package form

import "ockham-api/model"

type ServicePlanForm struct {
	PlanTitle       string // 标题
	PlanDescription string // 描述
	PlanEnabled     bool   // 启用中
	PlanPrice       float32

	ServingDays int16

	BundledTrafficPlanID  uint
	AvailableTrafficPlans []uint
}

func (_this *ServicePlanForm) ToModel() *model.ServicePlan {
	return &model.ServicePlan{
		PlanTitle:            _this.PlanTitle,
		PlanDescription:      _this.PlanDescription,
		PlanEnabled:          _this.PlanEnabled,
		PlanPrice:            _this.PlanPrice,
		ServingDays:          _this.ServingDays,
		BundledTrafficPlanID: _this.BundledTrafficPlanID,
	}
}
