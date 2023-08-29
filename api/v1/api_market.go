package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"ockham-api/api/v1/form"
	"ockham-api/api/v1/middleware"
	"ockham-api/api/v1/util"
	"ockham-api/database"
	"ockham-api/model"
	"strconv"
	"time"
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

// CreateServicePlan
// @Summary Create service plan
// @Description Create service plan
// @Tags market
// @Security Bearer
// @Param param body form.ServicePlanForm true "Create service plan form"
// @Success 200 {object} util.Pack
// @Failure 409,500 {object} util.Pack
// @Router /v1/service-plans [POST]
func CreateServicePlan(c *gin.Context) {
	servicePlanForm := &form.ServicePlanForm{}
	util.FillJsonForm(c, servicePlanForm)
	err := database.Create(c, servicePlanForm.ToModel(), "ServicePlan", util.ErrorMessageStatus)
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
	idStr := c.Param("service_plan_id")
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
	ServicePlanIDs           []uint `json:"service_plan_ids"`
	AdditionalTrafficPlanIDs []uint `json:"traffic_plan_ids"`
}

// SubscribeServicePlan
// @Summary Subscribes service plan
// @Description Subscribes service plan
// @Tags market
// @Security Bearer
// @Param param body SubscribeServicePlanForm true "Subscribes service plan form"
// @Success 200 {object} util.Pack
// @Failure 409,500 {object} util.Pack
// @Router /v1/users/me/service-plan-subscriptions [POST]
func SubscribeServicePlan(c *gin.Context) {
	currentUser := middleware.GetCurrentUser(c)
	subscribeServicePlanForm := &SubscribeServicePlanForm{}
	util.FillJsonForm(c, subscribeServicePlanForm)

	billing := &model.Billing{
		BillingTitle:       fmt.Sprintf("订阅服务计划"),
		BillingDescription: fmt.Sprintf("订阅服务计划"),
		BillingTotal:       0.0,
		BillingDate:        time.Now(),
		PaymentDueDate:     time.Now().AddDate(0, 0, 1),
		PaymentSettled:     false,
		User:               currentUser,
	}
	servicePlans := model.GetServicePlans(subscribeServicePlanForm.ServicePlanIDs)
	trafficPlans := model.GetTrafficPlans(subscribeServicePlanForm.AdditionalTrafficPlanIDs)

	billing.SubscribeServicePlans(servicePlans, c)
	billing.SubscribeTrafficPlans(trafficPlans, c)

	billing.Save(c)

	// TODO send billing email

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
