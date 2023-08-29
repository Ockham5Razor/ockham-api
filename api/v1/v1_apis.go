package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ockham-api/api/v1/middleware"
	"ockham-api/api/v1/util"
)

func ApiV1(r *gin.Engine) {
	v1Group := r.Group("/api/v1")
	{
		v1GroupAuth := v1Group.Group("/auth")
		{
			v1GroupAuth.POST("/users", CreateUser)
			v1GroupAuth.GET("/users/me", middleware.Token(), middleware.CurrentUser(), GetMe)
			v1GroupAuth.POST("/users/:user_id/roles", middleware.Token(), middleware.HasAnyRole("admin"), GrantRole)
			v1GroupAuth.DELETE("/users/:user_id/roles/:role_id", middleware.Token(), middleware.HasAnyRole("admin"), RevokeRole)
			v1GroupAuth.POST("/sessions", CreateSession)
			v1GroupAuth.PUT("/sessions/any/renewing", RenewSession)
			v1GroupAuth.PUT("/email-validations/any/validating", ValidateEmail)
		}
		v1GroupRechargeCode := v1Group.Group("/recharge-codes")
		{
			v1GroupRechargeCode.POST(
				"/none/batch-generations",
				middleware.Token(),
				middleware.HasAllRoles("admin"),
				GenerateRechargeCodesInBatches,
			)
		}
		v1GroupWallet := v1Group.Group("/")
		{
			v1GroupWallet.GET("/users/me/wallet", middleware.Token(), middleware.HasAnyRole("user"), GetWalletInfo)
			v1GroupWallet.PUT("/users/me/wallet/recharging", middleware.Token(), middleware.HasAnyRole("user"), RechargeWallet)
			v1GroupWallet.GET("/users/me/wallet/records", middleware.Token(), middleware.HasAnyRole("user"), GetRecordsOfWallet)
		}
		v1GroupServicePlans := v1Group.Group("/")
		{
			v1GroupServicePlans.GET("/service-plans/:service_plan_id", GetServicePlans)
			v1GroupServicePlans.GET("/service-plans", ListServicePlans)
			v1GroupServicePlans.POST("/service-plans", middleware.Token(), middleware.HasAnyRole("admin"), CreateServicePlan)
			v1GroupServicePlans.GET("/users/me/service-plan-subscriptions/:service_plan_subscription_id", middleware.Token(), middleware.HasAnyRole("user"), GetMyServicePlanSubscriptions)
			v1GroupServicePlans.GET("/users/me/service-plan-subscriptions", middleware.Token(), middleware.HasAnyRole("user"), ListMyServicePlanSubscriptions)
			v1GroupServicePlans.POST("/users/me/service-plan-subscriptions", middleware.Token(), middleware.HasAnyRole("user"), SubscribeServicePlan)
		}
		v1GroupAgents := v1Group.Group("/")
		{
			v1GroupAgents.GET("/agents/:agent_id/config", GetAgentConfig)
			v1GroupAgents.PUT("/agents/:agent_id/pulse", middleware.Signature(), middleware.SignatureCheck("agent_id", "agent_pulse", GetAgentSecretKey), AgentPulse)
		}
		v1GroupSubscriptions := v1Group.Group("/")
		{
			v1GroupSubscriptions.GET("/subscription-views/:id", ViewSubscription)
		}
	}
}

func DefaultHttp404(r *gin.Engine) {
	r.NoRoute(func(c *gin.Context) {
		util.ErrorMessageStatus(c, fmt.Sprintf("Not found for request [%v] %v", c.Request.Method, c.Request.URL.Path), 404)
	})
}
