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
				apiV1Util.ErrorPack(c).WithMessage(fmt.Sprintf("Token extracting failed: %s.", jwtClaimsError.Error())).WithHttpResponseCode(http.StatusBadRequest).Responds()
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
					apiV1Util.ErrorPack(c).WithMessage(fmt.Sprintf("User not authorized as %s", neededRoles[i])).WithHttpResponseCode(http.StatusForbidden).Responds()
					c.Abort()
					return
				}
			}
			c.Next()
			return
		} else {
			apiV1Util.ErrorPack(c).WithMessage("Token extracting failed, maybe you should use token middleware first.").WithHttpResponseCode(http.StatusForbidden).Responds()
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
				apiV1Util.ErrorPack(c).WithMessage(fmt.Sprintf("Token extracting failed: %s.", jwtClaimsError.Error())).WithHttpResponseCode(http.StatusBadRequest).Responds()
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
			apiV1Util.ErrorPack(c).WithMessage("User is not authorized as any required roles.").WithHttpResponseCode(http.StatusForbidden).Responds()
			c.Abort()
			return
		} else {
			apiV1Util.ErrorPack(c).WithMessage("Token extracting failed, maybe you should use token middleware first.").WithHttpResponseCode(http.StatusForbidden).Responds()
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
				apiV1Util.ErrorPack(c).WithMessage(fmt.Sprintf("Token extracting failed: %s.", jwtClaimsError.Error())).WithHttpResponseCode(http.StatusBadRequest).Responds()
				c.Abort()
				return
			}
			username := jwtClaims.Username
			user := &model.User{}
			database.DBConn.First(user, &model.User{Username: username})
			c.Set("user", user)
			c.Next()
		} else {
			apiV1Util.ErrorPack(c).WithMessage("Token extracting failed, maybe you should use token middleware first.").WithHttpResponseCode(http.StatusForbidden).Responds()
			c.Abort()
			return
		}
	}
}