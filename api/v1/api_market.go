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
	ServicePlans []struct {
		ServicePlanID          uint `json:"service_plan_id"`
		AdditionalTrafficPlans []struct {
			TrafficPlanID uint `json:"traffic_plan_id"`
		} `json:"additional_traffic_plans"`
	} `json:"service_plans"`
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
	f := &SubscribeServicePlanForm{}
	util.FillJsonForm(c, f)

	now := time.Now()

	billing := &model.Billing{
		BillingTitle:       fmt.Sprintf("订阅服务计划"),
		BillingDescription: fmt.Sprintf("订阅服务计划"),
		BillingTotal:       0.0,
		BillingDate:        now,
		PaymentDueDate:     now.AddDate(0, 0, 1),
		PaymentSettled:     false,
		User:               currentUser,
	}

	spSubIDs := make([]uint, 0)
	tpSubIDs := make([]uint, 0)
	for _, spReq := range f.ServicePlans {
		// create service plan
		sp := model.GetServicePlan[model.ServicePlan](spReq.ServicePlanID)
		spSub := &model.ServicePlanSubscription{
			SubscriptionTitle:       sp.PlanTitle,
			SubscriptionDescription: sp.PlanDescription,
			SubscriptionEnabled:     false,
			SubscriptionStartTime:   now,
			SubscriptionEndTime:     now.AddDate(0, 0, sp.ServingDays),
			ServicePlanID:           sp.ID,
			UserID:                  currentUser.ID,
		}
		_ = database.Create(c, spSub, "ServicePlanSubscription", util.ErrorMessageStatus)
		spSubIDs = append(spSubIDs, spSub.ID)
		billing.BillingTotal += sp.PlanPrice

		// create bundled traffic plan
		btpSub := &model.TrafficPlanSubscription{
			SubscriptionTitle:         sp.BundledTrafficPlan.PlanTitle,
			SubscriptionDescription:   sp.BundledTrafficPlan.PlanDescription,
			SubscriptionEnabled:       false,
			SubscriptionStartTime:     now,
			SubscriptionEndTime:       now.AddDate(0, 0, sp.ServingDays),
			TotalTrafficBytes:         sp.BundledTrafficPlan.TotalTrafficBytes,
			UsedTrafficBytes:          0,
			SystemPriority:            0,
			UserPriority:              0,
			AdminPriority:             0,
			ServicePlanID:             sp.ID,
			ServicePlanSubscriptionID: spSub.ID,
			UserID:                    currentUser.ID,
			Bundled:                   true,
		}
		_ = database.Create(c, btpSub, "TrafficPlanSubscription", util.ErrorMessageStatus)
		tpSubIDs = append(tpSubIDs, btpSub.ID)

		// create additional traffic plan
		priorityRank := 100
		for _, tpReq := range spReq.AdditionalTrafficPlans {
			priorityRank -= 1 // priority decrease
			tp := model.GetTrafficPlan[model.TrafficPlan](tpReq.TrafficPlanID)
			tpSub := &model.TrafficPlanSubscription{
				SubscriptionTitle:         tp.PlanTitle,
				SubscriptionDescription:   tp.PlanDescription,
				SubscriptionEnabled:       false,
				SubscriptionStartTime:     now,
				SubscriptionEndTime:       now.AddDate(0, 0, sp.ServingDays),
				TotalTrafficBytes:         tp.TotalTrafficBytes,
				UsedTrafficBytes:          0,
				SystemPriority:            priorityRank,
				UserPriority:              0,
				AdminPriority:             0,
				ServicePlanID:             sp.ID,
				ServicePlanSubscriptionID: spSub.ID,
				UserID:                    currentUser.ID,
				Bundled:                   false,
			}
			_ = database.Create(c, tpSub, "TrafficPlanSubscription", util.ErrorMessageStatus)
			tpSubIDs = append(tpSubIDs, tpSub.ID)
			billing.BillingTotal += tp.PlanPrice
		}
	}

	billing.SubscribingServicePlans = spSubIDs
	billing.SubscribingTrafficPlans = tpSubIDs
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
