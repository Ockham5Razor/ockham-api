package model

import "gorm.io/gorm"

type ServicePlanUtilization struct {
	gorm.Model

	Title                 string  `gorm:"type:VARCHAR(24)"` // 标题
	Description           string  `gorm:"type:LONGTEXT"`    // 描述
	CyclicalTrafficBytes  int64   // 每次循环的流量大小
	CyclicalIntervalDays  int16   // 循环周期
	InheritSurplusTraffic bool    // 循环中继承结余流量
	TotalCycleTimes       int16   // 总循环次数
	CycleTimes            int16   // 已循环次数
	Fee                   float32 // 费用
	UserID                uint
	User                  User
}
