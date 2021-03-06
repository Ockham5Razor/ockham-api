package model

import (
	"gorm.io/gorm"
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

	ServicePlanSubscriptionID uint                     // 用户的订阅（外键）
	ServicePlanSubscription   *ServicePlanSubscription // 用户的订阅（引用）
	UserID                    uint                     // 用户（外键）
	User                      *User                    // 用户（引用）
}
