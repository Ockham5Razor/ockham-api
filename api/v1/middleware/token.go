package middleware

import (
	"github.com/gin-gonic/gin"
	"gol-c/api/v1/util"
	"gol-c/database"
	"gol-c/model"
	"net/http"
	"strings"
)

func Token() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			util.ErrorMessageStatus(c, "HTTP header `Authorization` is required for this request.", http.StatusBadRequest)
			c.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			util.ErrorMessageStatus(c, "HTTP header `Authorization` is in illegal format.", http.StatusBadRequest)
			c.Abort()
			return
		}

		session := &model.Session{}
		database.GetByField(&model.Session{SessionBody: parts[1]}, session, []string{"User"})
		c.Set("session", session)
		c.Next()
	}
}
