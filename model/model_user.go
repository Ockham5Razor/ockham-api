package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"ockham-api/api/v1/util"
	"ockham-api/database"
	"time"
)

type User struct {
	gorm.Model    `json:"-"`
	Username      string  `gorm:"type:VARCHAR(24);uniqueIndex" json:"username"`
	Password      string  `gorm:"type:VARCHAR(128)" json:"-"`
	Email         string  `gorm:"type:VARCHAR(128)" json:"email"`
	EmailVerified bool    `json:"email_verified"`
	Roles         []*Role `gorm:"many2many:user_role;" json:"-"`
}

type Role struct {
	gorm.Model
	RoleName string `gorm:"type:VARCHAR(24)"`
}

func GetUser(userID uint64) *User {
	user := &User{}
	database.DBConn.Preload("Roles").First(user, userID)
	return user
}

func (user *User) Subscribes(sp *ServicePlan, c *gin.Context) {
	subscription := &ServicePlanSubscription{
		SubscriptionTitle:       sp.PlanTitle,
		SubscriptionDescription: sp.PlanDescription,
		SubscriptionEnabled:     false,

		TermSpacingDays: sp.TermSpacingDays,
		TermsRemaining:  sp.TermsTotal,
		TermLastDate:    time.Now(),

		EachTermIncreasingTrafficBytes: sp.EachTermIncreasingTrafficBytes,
		EachTermInheritsSurplusTraffic: sp.EachTermInheritsSurplusTraffic,

		SubscriptionRenewalAddingTerms:    sp.TermsTotal,
		SubscriptionRenewalTimesRemaining: 0, // 续期次数：0，不可续订，只能订阅新 plan。
		SubscriptionRenewalFee:            sp.SubscriptionFee,

		User: user,
	}
	_ = database.Create(c, subscription, "ServicePlanSubscription", util.ErrorMessageStatus)

	billing := &Billing{
		BillingTitle:            fmt.Sprintf("订阅服务计划"),
		BillingDescription:      sp.PlanDescription,
		BillingTotal:            sp.SubscriptionFee,
		BillingDate:             time.Now(),
		PaymentDueDate:          time.Now().AddDate(0, 0, 1),
		PaymentSettled:          false,
		ServicePlanSubscription: subscription,
		User:                    user,
	}
	_ = database.Create(c, billing, "Billing", util.ErrorMessageStatus)
}

func (user *User) RemoveRole(roleId uint) {
	targetIndex := -1
	for i, role := range user.Roles {
		if role.ID == roleId {
			fmt.Println(i)
			targetIndex = i
		}
	}
	fmt.Println(user.Roles)
	if targetIndex != -1 {
		user.Roles = append(user.Roles[:targetIndex], user.Roles[targetIndex+1:]...)
	}
	fmt.Println(user.Roles)
	_ = database.DBConn.Model(user).Association("Roles").Replace(user.Roles)
}
