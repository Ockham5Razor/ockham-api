package model

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"ockham-api/api/v1/util"
	"ockham-api/database"
	"time"
)

type Billing struct {
	gorm.Model

	BillingTitle       string    `gorm:"type:VARCHAR(24)"` // 账单标题
	BillingDescription string    `gorm:"type:LONGTEXT"`    // 账单备注
	BillingTotal       float32   // 账单总额
	BillingDate        time.Time // 账单日期：用于展示。

	PaymentSettled   bool      // 是否已付清
	PaymentStartDate time.Time // 付款开始日期：此前付款无效，账单将提前此日期 7 天展示。
	PaymentDueDate   time.Time // 付款截止日期：需在此前付款，付款截止后将停止服务。

	SubscribingServicePlans []uint // 订阅服务
	SubscribingTrafficPlans []uint // 订阅流量（包含内置流量）

	UserID uint  // 用户（外键）
	User   *User // 用户（引用）
}

func (b *Billing) Save(c *gin.Context) {
	_ = database.Create(c, b, "Billing", util.ErrorMessageStatus)
}

func (b *Billing) AllSubscriptionActivate() {
	// activate service plan subscription
	spSubs := make([]ServicePlanSubscription, 0)
	database.GetMore[ServicePlanSubscription](b.SubscribingServicePlans, &spSubs).Update("SubscriptionEnabled", true)

	// activate additional traffic plan subscription
	tpSubs := make([]TrafficPlanSubscription, 0)
	database.GetMore[TrafficPlanSubscription](b.SubscribingTrafficPlans, &tpSubs).Update("SubscriptionEnabled", true)

	// create traffic packs
	for _, tpSub := range tpSubs {
		t := TrafficPack{
			TotalTrafficBytes: tpSub.Traffic,
			UsedTrafficBytes:  0,

			SystemPriority: tpSub.SystemPriority,
			UserPriority:   tpSub.UserPriority,
			AdminPriority:  tpSub.AdminPriority,

			ServicePlanSubscriptionID: tpSub.ServicePlanSubscriptionID,
			TrafficPlanSubscriptionID: tpSub.ID,

			UserID: tpSub.UserID,
		}
		database.DBConn.Save(t)
	}
}
