package v1

import (
	"github.com/gin-gonic/gin"
	"gol-c/api/v1/middleware"
)

func ApiV1(r *gin.Engine) {
	v1Group := r.Group("/api/v1")
	{
		v1GroupAuth := v1Group.Group("/auth")
		{
			v1GroupAuth.POST("/users", CreateUser)
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
			v1GroupServicePlans.GET("/service-plans", ListServicePlans)
			v1GroupServicePlans.POST("/service-plans", middleware.Token(), middleware.HasAnyRole("admin"), CreateServicePlan)
			v1GroupServicePlans.GET("/users/me/service-plan-subscriptions", middleware.Token(), middleware.HasAnyRole("user"), ListMyServicePlanSubscriptions)
			v1GroupServicePlans.POST("/users/me/service-plan-subscriptions", middleware.Token(), middleware.HasAnyRole("user"), SubscribeServicePlan)
		}
	}
}
