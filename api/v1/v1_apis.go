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
			v1GroupAuth.PUT("/sessions/any:renew", RenewSession)
			v1GroupAuth.PUT("/email-validations/any:validate", ValidateEmail)
		}
		v1GroupRechargeCode := v1Group.Group("/recharge-codes")
		{
			v1GroupRechargeCode.POST(
				"/none/batch-generations",
				middleware.Token(),
				middleware.HasAllAuthorities("admin"),
				GenerateRechargeCodesInBatches,
			)
		}
	}
}
