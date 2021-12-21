package model

import (
	"gol-c/database"
	"gorm.io/gorm"
)

type ServicePlan struct {
	gorm.Model

	PlanTitle       string `gorm:"type:VARCHAR(24)"` // 标题
	PlanDescription string `gorm:"type:LONGTEXT"`    // 描述
	PlanEnabled     bool   // 启用中

	TermSpacingDays int16 // 用期周期
	TermsTotal      int16 // 用期总数

	EachTermIncreasingTrafficBytes int64 // 每次用期增加的流量大小
	EachTermInheritsSurplusTraffic bool  // 每次用期循环继承结余流量

	SubscriptionFee float32 // 订阅 / 更新单价
}

func GetServicePlan(id uint) *ServicePlan {
	servicePlan := &ServicePlan{}
	database.Get(id, servicePlan)
	return servicePlan
}
