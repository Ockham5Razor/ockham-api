package model

import (
	"gorm.io/gorm"
)

type ServicePlan struct {
	gorm.Model

	Title                 string  `gorm:"type:VARCHAR(24)"` // 标题
	Description           string  `gorm:"type:LONGTEXT"`    // 描述
	CyclicalTrafficBytes  int64   // 每次循环的流量大小
	CyclicalLastingDays   int16   // 循环周期
	InheritSurplusTraffic bool    // 循环中继承结余流量
	TotalCycleTimes       int16   // 总循环次数
	PriceForEachCycle     float32 // 每次循环价格
	Enabled               bool    // 启用中
}

func (_this *ServicePlan) Subscribe(user *User) *ServicePlanSubscription {
	return &ServicePlanSubscription{
		Title:                 _this.Title,
		Description:           _this.Description,
		CyclicalTrafficBytes:  _this.CyclicalTrafficBytes,
		CyclicalLastingDays:   _this.CyclicalLastingDays,
		InheritSurplusTraffic: _this.InheritSurplusTraffic,
		TotalCycleTimes:       _this.TotalCycleTimes,
		FeeForEachCycle:       _this.PriceForEachCycle,
		User:                  *user,
		Enabled:               false,
	}
}
