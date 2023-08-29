package v1

import "github.com/gin-gonic/gin"

// ViewSubscription
// @Summary ViewSubscription
// @Description Subscribe your nodes by link
// @Tags auth
// @Success 201 {object} util.Pack
// @Failure 401,409,500 {object} util.Pack
// @Param client param string true "Client Type"
// @Param token param string true "Subscription Token"
// @Param id path uint true "Subscription Token"
// @Router /v1/subscriptions/{id}/subscribe [POST]
func ViewSubscription(c *gin.Context) {
	_ = c.Param("id")
}

func ListSubscriptions(c *gin.Context) {
}

func GetSubscription(c *gin.Context) {
}
