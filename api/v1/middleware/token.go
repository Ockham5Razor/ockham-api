package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ockham-api/api/v1/util"
	"strings"
)

func Token() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			util.ErrorPack(c).WithMessage("HTTP header `Authorization` is required for this request.").WithHttpResponseCode(http.StatusBadRequest).Responds()
			c.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			util.ErrorPack(c).WithMessage("HTTP header `Authorization` is in illegal format.").WithHttpResponseCode(http.StatusBadRequest).Responds()
			c.Abort()
			return
		}
		token := parts[1]
		c.Set("token", token)
		c.Next()
	}
}
