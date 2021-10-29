package v1

import (
	"github.com/gin-gonic/gin"
	"gol-c/api/v1/util"
	"gol-c/database"
	"gol-c/model"
)

// ListServicePlans
// @Summary Get all service plans
// @Description Get all service plans
// @Tags market
// @Success 200 {object} util.Pack
// @Router /v1/service-plans [GET]
func ListServicePlans(c *gin.Context) {
	servicePlans := &[]model.ServicePlan{}
	database.DBConn.Find(servicePlans)
	util.SuccessPack(c).WithData(servicePlans).Responds()
}

type ServicePlanForm struct {
	Title                 string  // 标题
	Description           string  // 描述
	CyclicalTrafficBytes  int64   // 每次循环的流量大小
	CyclicalIntervalDays  int16   // 循环周期
	InheritSurplusTraffic bool    // 循环中继承结余流量
	TotalCycleTimes       int16   // 总循环次数
	Price                 float32 // 价格
}

func (servicePlanForm *ServicePlanForm) toModel() *model.ServicePlan {
	return &model.ServicePlan{
		Title:                 servicePlanForm.Title,
		Description:           servicePlanForm.Description,
		CyclicalTrafficBytes:  servicePlanForm.CyclicalTrafficBytes,
		CyclicalIntervalDays:  servicePlanForm.CyclicalIntervalDays,
		InheritSurplusTraffic: servicePlanForm.InheritSurplusTraffic,
		TotalCycleTimes:       servicePlanForm.TotalCycleTimes,
		Price:                 servicePlanForm.Price,
	}
}

// CreateServicePlan
// @Summary Create service plan
// @Description Create service plan
// @Tags market
// @Security Bearer
// @Param param body ServicePlanForm true "Create service plan form"
// @Success 200 {object} util.Pack
// @Failure 409,500 {object} util.Pack
// @Router /v1/service-plans [POST]
func CreateServicePlan(c *gin.Context) {
	servicePlanForm := &ServicePlanForm{}
	util.GetJsonForm(c, servicePlanForm)
	err := database.Create(c, servicePlanForm.toModel(), "ServicePlan", util.ErrorMessageStatus)
	if err != nil {
		return
	}
	util.SuccessPack(c).WithMessage("Service plan created!").Responds()
}
