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

	SystemPriority uint // 系统优先级
	UserPriority   uint // 用户优先级

	ServicePlanID uint        `json:"-"`
	ServicePlan   ServicePlan `json:"-"`

	ServicePlanSubscriptionID uint                    `json:"-"`
	ServicePlanSubscription   ServicePlanSubscription `json:"-"`

	UserID uint `json:"-"`
	User   User `json:"-"`
}

type TrafficPacks struct {
	gorm.Model

	TotalTrafficBytes uint64
	UsedTrafficBytes  uint64

	StartTime time.Time
	EndTime   time.Time

	SystemPriority uint // 系统优先级
	UserPriority   uint // 用户优先级

	ServicePlanSubscriptionID uint                    `json:"-"`
	ServicePlanSubscription   ServicePlanSubscription `json:"-"`

	TrafficPlanSubscriptionID uint                    `json:"-"`
	TrafficPlanSubscription   TrafficPlanSubscription `json:"-"`

	UserID uint `json:"-"`
	User   User `json:"-"`
}
