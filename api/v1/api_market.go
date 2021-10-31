package v1

import (
	"github.com/gin-gonic/gin"
	"gol-c/api/v1/middleware"
	"gol-c/api/v1/util"
	"gol-c/database"
	"gol-c/model"
	"net/http"
	"strconv"
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
	CyclicalLastingDays   int16   // 循环周期
	InheritSurplusTraffic bool    // 循环中继承结余流量
	TotalCycleTimes       int16   // 总循环次数
	PriceForEachCycle     float32 // 每次循环价格
	Enabled               bool    // 启用中
}

func (_this *ServicePlanForm) toModel() *model.ServicePlan {
	return &model.ServicePlan{
		Title:                 _this.Title,
		Description:           _this.Description,
		CyclicalTrafficBytes:  _this.CyclicalTrafficBytes,
		CyclicalLastingDays:   _this.CyclicalLastingDays,
		InheritSurplusTraffic: _this.InheritSurplusTraffic,
		TotalCycleTimes:       _this.TotalCycleTimes,
		PriceForEachCycle:     _this.PriceForEachCycle,
		Enabled:               _this.Enabled,
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

// GetServicePlans
// @Summary Get service plan
// @Description Get service plan
// @Tags market
// @Param service_plan_id path int true "service plan id"
// @Success 200 {object} util.Pack
// @Router /v1/service-plans/{service_plan_id} [GET]
func GetServicePlans(c *gin.Context) {
	idStr := c.Param("service_plan_subscription_id")
	idU64, _ := strconv.ParseUint(idStr, 10, 32)
	servicePlan := &model.ServicePlan{}
	ctx := database.DBConn.Find(servicePlan, idU64)
	if ctx.RowsAffected == 0 {
		util.ErrorPack(c).WithHttpResponseCode(http.StatusNotFound).WithMessage("Service plan not found!").Responds()
	} else {
		util.SuccessPack(c).WithData(servicePlan).Responds()
	}
}

// ListMyServicePlanSubscriptions
// @Summary List service plan subscriptions
// @Description List service plan subscriptions
// @Tags market
// @Security Bearer
// @Success 200 {object} util.Pack
// @Failure 409,500 {object} util.Pack
// @Router /v1/users/me/service-plan-subscriptions [GET]
func ListMyServicePlanSubscriptions(c *gin.Context) {
	currentUser := middleware.GetCurrentUser(c)

	subscriptions := &[]model.ServicePlanSubscription{}
	database.GetByField(model.ServicePlanSubscription{UserID: currentUser.ID}, subscriptions, nil)

	util.SuccessPack(c).WithData(subscriptions).Responds()
}

type SubscribeServicePlanForm struct {
	ServicePlanId       uint
	ConsolidateBillings bool // 合并账单一次结清
}

// SubscribeServicePlan
// @Summary Subscribe service plan
// @Description Subscribe service plan
// @Tags market
// @Security Bearer
// @Param param body SubscribeServicePlanForm true "Subscribe service plan form"
// @Success 200 {object} util.Pack
// @Failure 409,500 {object} util.Pack
// @Router /v1/users/me/service-plan-subscriptions [POST]
func SubscribeServicePlan(c *gin.Context) {
	currentUser := middleware.GetCurrentUser(c)
	subscribeServicePlanForm := &SubscribeServicePlanForm{}
	util.GetJsonForm(c, subscribeServicePlanForm)

	servicePlan := &model.ServicePlan{}
	database.Get(subscribeServicePlanForm.ServicePlanId, servicePlan)

	servicePlanSubscription := servicePlan.Subscribe(currentUser)
	err := database.Create(c, servicePlanSubscription, "ServicePlanSubscription", util.ErrorMessageStatus)
	if err != nil {
		return
	}
	util.SuccessPack(c).WithMessage("Successfully subscribed service plan!").Responds()
}

// GetMyServicePlanSubscriptions
// @Summary Get service plan subscriptions
// @Description Get service plan subscriptions
// @Tags market
// @Security Bearer
// @Param service_plan_subscription_id path int true "service plan subscriptions id"
// @Success 200 {object} util.Pack
// @Failure 409,500 {object} util.Pack
// @Router /v1/users/me/service-plan-subscriptions/{service_plan_subscription_id} [GET]
func GetMyServicePlanSubscriptions(c *gin.Context) {
	currentUser := middleware.GetCurrentUser(c)
	idStr := c.Param("service_plan_subscription_id")
	idU64, _ := strconv.ParseUint(idStr, 10, 32)
	subscription := &model.ServicePlanSubscription{}
	ctx := database.DBConn.Where(model.ServicePlanSubscription{UserID: currentUser.ID}).Find(subscription, idU64)
	if ctx.RowsAffected == 0 {
		util.ErrorPack(c).WithHttpResponseCode(http.StatusNotFound).WithMessage("Service plan subscription not found!").Responds()
	} else {
		util.SuccessPack(c).WithData(subscription).Responds()
	}
}
