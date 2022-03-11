package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	apiV1Util "ockham-api/api/v1/util"
	"ockham-api/database"
	"ockham-api/model"
	"ockham-api/util"
	"strings"
)

func HasAllRoles(neededRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, tokenExists := c.Get(apiV1Util.ContextBearerValue)
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
			c.Set(apiV1Util.ContextCurrentUser, user)

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
			apiV1Util.ErrorPack(c).WithMessage("Token extracting failed, maybe you should use token middleware first.").WithHttpResponseCode(http.StatusInternalServerError).Responds()
			c.Abort()
			return
		}
	}
}

func HasAnyRole(neededRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, tokenExists := c.Get(apiV1Util.ContextBearerValue)
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
			c.Set(apiV1Util.ContextCurrentUser, user)

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
			apiV1Util.ErrorPack(c).WithMessage("Token extracting failed, maybe you should use token middleware first.").WithHttpResponseCode(http.StatusInternalServerError).Responds()
			c.Abort()
			return
		}
	}
}

func CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, tokenExists := c.Get(apiV1Util.ContextBearerValue)
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
			c.Set(apiV1Util.ContextCurrentUser, user)
			c.Next()
		} else {
			apiV1Util.ErrorPack(c).WithMessage("Token extracting failed, maybe you should use token middleware first.").WithHttpResponseCode(http.StatusForbidden).Responds()
			c.Abort()
			return
		}
	}
}

func GetCurrentUser(c *gin.Context) *model.User {
	userIntf, userExists := c.Get(apiV1Util.ContextCurrentUser)
	if userExists {
		user := userIntf.(*model.User)
		return user
	} else {
		apiV1Util.ErrorPack(c).WithMessage("Fatal: extracting current user failed.").WithHttpResponseCode(http.StatusInternalServerError).Responds()
		c.Abort()
		return nil
	}
}

func SignatureCheck(resourceIdPathParamName string, checkFunc func(AccessKeyID string, signatureStrFromClient string, timestampStr string) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		resourceIdStr := c.Param(resourceIdPathParamName)
		signatureValueObj, signatureExists := c.Get(apiV1Util.ContextSignatureValue)
		timestampValueObj, timestampExists := c.Get(apiV1Util.ContextSignatureValue)
		if signatureExists && timestampExists {
			err := checkFunc(resourceIdStr, signatureValueObj.(string), timestampValueObj.(string))
			if err != nil {
				apiV1Util.ErrorPack(c).WithMessage(err.Error()).WithHttpResponseCode(http.StatusUnauthorized).Responds()
				c.Abort()
			} else {
				c.Next()
			}
		} else {
			apiV1Util.ErrorPack(c).WithMessage("Signature extracting failed, maybe you should use signature middleware first.").WithHttpResponseCode(http.StatusForbidden).Responds()
			c.Abort()
		}
	}
}
