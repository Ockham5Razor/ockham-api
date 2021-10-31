package model

import "gorm.io/gorm"

type ServicePlanSubscription struct {
	gorm.Model

	Title                 string  `gorm:"type:VARCHAR(24)"` // 标题
	Description           string  `gorm:"type:LONGTEXT"`    // 描述
	CyclicalTrafficBytes  int64   // 每次循环的流量大小
	CyclicalLastingDays   int16   // 循环周期
	InheritSurplusTraffic bool    // 循环中继承结余流量
	TotalCycleTimes       int16   // 总循环次数，-1 为不限次数，如果为 -1 则不允许合并账单支付
	FeeForEachCycle       float32 // 每次循环费用
	UserID                uint
	User                  User `json:"-"`
	Enabled               bool // 启用
}
