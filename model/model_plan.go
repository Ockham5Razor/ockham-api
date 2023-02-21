package model

import (
	"gorm.io/gorm"
	"ockham-api/database"
)

type ServicePlan struct {
	gorm.Model

	PlanTitle       string `gorm:"type:VARCHAR(24)"`
	PlanDescription string `gorm:"type:LONGTEXT"`
	PlanEnabled     bool
	PlanPrice       float32

	ServingDays int16

	BundledTrafficPlanID uint
	BundledTrafficPlan   TrafficPlan

	AvailableTrafficPlans []*TrafficPlan `gorm:"many2many:service_plan_available_traffic_plan;"`
}

type TrafficPlan struct {
	gorm.Model

	PlanTitle       string `gorm:"type:VARCHAR(24)"`
	PlanDescription string `gorm:"type:LONGTEXT"`
	PlanEnabled     bool
	PlanPrice       float32

	MonthlyTrafficBytes int64
	Inheritable         bool
}

func GetServicePlan(id uint) *ServicePlan {
	servicePlan := &ServicePlan{}
	database.Get(id, servicePlan)
	return servicePlan
}
