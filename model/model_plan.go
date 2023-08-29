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

	ServingDays int

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
	plan := &ServicePlan{}
	database.Get[ServicePlan](id, plan)
	return plan
}

func GetServicePlans(ids []uint) *[]ServicePlan {
	plans := &[]ServicePlan{}
	database.GetMore[ServicePlan](ids, plans)
	return plans
}

func GetTrafficPlan(id uint) *TrafficPlan {
	plan := &TrafficPlan{}
	database.Get[TrafficPlan](id, plan)
	return plan
}

func GetTrafficPlans(ids []uint) *[]TrafficPlan {
	plans := &[]TrafficPlan{}
	database.GetMore[TrafficPlan](ids, plans)
	return plans
}
