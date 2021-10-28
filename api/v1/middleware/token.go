package middleware

import (
	"github.com/gin-gonic/gin"
	"gol-c/api/v1/util"
	"net/http"
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
