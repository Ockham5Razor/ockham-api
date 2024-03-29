package model

import (
	"gorm.io/gorm"
	"time"
)

type ServicePlanSubscription struct {
	gorm.Model

	SubscriptionTitle       string `gorm:"type:VARCHAR(24)"`
	SubscriptionDescription string `gorm:"type:LONGTEXT"`
	SubscriptionEnabled     bool
	SubscriptionStartTime   time.Time
	SubscriptionEndTime     time.Time

	ServicePlanID uint        `json:"-"`
	ServicePlan   ServicePlan `json:"-"`

	UserID uint `json:"-"`
	User   User `json:"-"`
}

type TrafficPlanSubscription struct {
	gorm.Model

	SubscriptionTitle       string `gorm:"type:VARCHAR(24)"`
	SubscriptionDescription string `gorm:"type:LONGTEXT"`
	SubscriptionEnabled     bool
	SubscriptionStartTime   time.Time
	SubscriptionEndTime     time.Time

	TotalTrafficBytes int64
	UsedTrafficBytes  int64

	// 优先级：越大越优先，排序依据：admin -> user -> system
	SystemPriority int // 系统预设的优先级
	UserPriority   int // 用户排列的优先级
	AdminPriority  int // 管理员排列的优先级

	ServicePlanID uint        `json:"-"`
	ServicePlan   ServicePlan `json:"-"`

	ServicePlanSubscriptionID uint                    `json:"-"`
	ServicePlanSubscription   ServicePlanSubscription `json:"-"`

	UserID uint `json:"-"`
	User   User `json:"-"`

	Bundled bool
}
