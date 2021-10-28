package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	apiV1Util "gol-c/api/v1/util"
	"gol-c/database"
	"gol-c/model"
	"gol-c/util"
	"net/http"
	"strings"
)

func HasAllRoles(neededRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, tokenExists := c.Get("token")
		if tokenExists {
			jwtClaims, jwtClaimsError := util.ParseToken(token.(string))
			if jwtClaimsError != nil {
				apiV1Util.ErrorMessageStatus(c, fmt.Sprintf("Token extracting failed: %s.", jwtClaimsError.Error()), http.StatusBadRequest)
				c.Abort()
				return
			}
			username := jwtClaims.Username
			user := &model.User{}
			database.DBConn.Preload("Roles").First(user, &model.User{Username: username})
			c.Set("user", user)

			hasRoleSet := make(map[string]bool)
			roles := user.Roles
			for i := range roles {
				hasRoleSet[strings.ToLower(roles[i].RoleName)] = true
			}
			for i := range neededRoles {
				_, hasOneNeededRole := hasRoleSet[strings.ToLower(neededRoles[i])]
				if !hasOneNeededRole {
					apiV1Util.ErrorMessageStatus(c, "User not authorized as "+neededRoles[i], http.StatusForbidden)
					c.Abort()
					return
				}
			}
			c.Next()
			return
		} else {
			apiV1Util.ErrorMessageStatus(c, "Token extracting failed, maybe you should use token middleware first.", http.StatusBadRequest)
			c.Abort()
			return
		}
	}
}

func HasAnyRole(neededRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, tokenExists := c.Get("token")
		if tokenExists {
			jwtClaims, jwtClaimsError := util.ParseToken(token.(string))
			if jwtClaimsError != nil {
				apiV1Util.ErrorMessageStatus(c, fmt.Sprintf("Token extracting failed: %s.", jwtClaimsError.Error()), http.StatusBadRequest)
				c.Abort()
				return
			}
			username := jwtClaims.Username
			user := &model.User{}
			database.DBConn.Preload("Roles").First(user, &model.User{Username: username})
			c.Set("user", user)

			hasRoleSet := make(map[string]bool)
			roles := user.Roles
			for i := range roles {
				hasRoleSet[strings.ToLower(roles[i].RoleName)] = true
			}
			for i := range neededRoles {
				_, hasOneNeededRole := hasRoleSet[strings.ToLower(neededRoles[i])]
				if hasOneNeededRole {
					c.Next()
					return
				}
			}
			apiV1Util.ErrorMessageStatus(c, "User is not authorized as any required roles.", http.StatusForbidden)
			c.Abort()
			return
		} else {
			apiV1Util.ErrorMessageStatus(c, "Token extracting failed, maybe you should use token middleware first.", http.StatusBadRequest)
			c.Abort()
			return
		}
	}
}

func CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, tokenExists := c.Get("token")
		if tokenExists {
			jwtClaims, jwtClaimsError := util.ParseToken(token.(string))
			if jwtClaimsError != nil {
				apiV1Util.ErrorMessageStatus(c, fmt.Sprintf("Token extracting failed: %s.", jwtClaimsError.Error()), http.StatusBadRequest)
				c.Abort()
				return
			}
			username := jwtClaims.Username
			user := &model.User{}
			database.DBConn.First(user, &model.User{Username: username})
			c.Set("user", user)
			c.Next()
		} else {
			apiV1Util.ErrorMessageStatus(c, "Token extracting failed, maybe you should use token middleware first.", http.StatusBadRequest)
			c.Abort()
			return
		}
	}
}
