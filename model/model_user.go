package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gol-c/api/v1/util"
	"gol-c/database"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Username      string `gorm:"type:VARCHAR(24);uniqueIndex"`
	Password      string `gorm:"type:VARCHAR(128)"`
	Email         string `gorm:"type:VARCHAR(128)"`
	EmailVerified bool
	Roles         []Role `gorm:"many2many:user_role;"`
}

type Role struct {
	gorm.Model
	RoleName string `gorm:"type:VARCHAR(24)"`
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
