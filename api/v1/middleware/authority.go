package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	apiV1Util "ockham-api/api/v1/util"
	"ockham-api/config"
	"ockham-api/database"
	"ockham-api/model"
	"ockham-api/util"
	"strconv"
	"strings"
	"time"
)

func HasAllRoles(neededRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, tokenExists := c.Get(apiV1Util.ContextBearerBody)
		if tokenExists {
			jwtClaims, jwtClaimsError := util.ParseToken(token.(string))
			if jwtClaimsError != nil {
				apiV1Util.ErrorPack(c).WithMessage("Token extracting failed: %s.", jwtClaimsError.Error()).WithHttpResponseCode(http.StatusBadRequest).Responds()
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
					apiV1Util.ErrorPack(c).WithMessage("User not authorized as %s", neededRoles[i]).WithHttpResponseCode(http.StatusForbidden).Responds()
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
		token, tokenExists := c.Get(apiV1Util.ContextBearerBody)
		if tokenExists {
			jwtClaims, jwtClaimsError := util.ParseToken(token.(string))
			if jwtClaimsError != nil {
				apiV1Util.ErrorPack(c).WithMessage("Token extracting failed: %s.", jwtClaimsError.Error()).WithHttpResponseCode(http.StatusBadRequest).Responds()
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
		token, tokenExists := c.Get(apiV1Util.ContextBearerBody)
		if tokenExists {
			jwtClaims, jwtClaimsError := util.ParseToken(token.(string))
			if jwtClaimsError != nil {
				apiV1Util.ErrorPack(c).WithMessage("Token extracting failed: %s.", jwtClaimsError.Error()).WithHttpResponseCode(http.StatusBadRequest).Responds()
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

func SignatureCheck(resourceIdPathParamName, actionType string, getResourceSecretFunc func(AccessKeyID string) (string, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		resourceIdStr := c.Param(resourceIdPathParamName)
		signatureValueObj, signatureExists := c.Get(apiV1Util.ContextSignatureValue)
		timestampValueObj, timestampExists := c.Get(apiV1Util.ContextTimestampValue)
		if signatureExists && timestampExists {
			// 校验时间戳
			timestampStr := timestampValueObj.(string)
			i, err := strconv.ParseInt(timestampStr, 10, 64)
			if err != nil {
				panic(err)
			}
			sigTime := time.Unix(i, 0)
			now := time.Now()
			secondsAgo := now.Add(-config.GetConfig().Auth.Signature.TimestampToleranceSeconds * time.Second)
			secondsAfter := now.Add(config.GetConfig().Auth.Signature.TimestampToleranceSeconds * time.Second)
			if sigTime.After(secondsAfter) || sigTime.Before(secondsAgo) { // 在容忍时间内
				apiV1Util.ErrorPack(c).WithMessage("expired or invalid signing time").WithHttpResponseCode(http.StatusBadRequest).Responds()
				c.Abort()
				return
			}

			// 取得资源密钥
			resourceSecretKey, err := getResourceSecretFunc(resourceIdStr)
			if err != nil {
				apiV1Util.ErrorPack(c).WithMessage(err.Error()).WithHttpResponseCode(http.StatusUnauthorized).Responds()
				c.Abort()
				return
			} else {
				err, sigFromClient := apiV1Util.DecodeSignature(signatureValueObj.(string))
				if err != nil {
					apiV1Util.ErrorPack(c).WithMessage(err.Error()).WithHttpResponseCode(http.StatusUnauthorized).Responds()
					c.Abort()
					return
				}
				sigGen, err := apiV1Util.CreateSignature(c.Request, resourceIdStr, resourceSecretKey, actionType)
				if err != nil {
					apiV1Util.ErrorPack(c).WithMessage("internal server error").WithHttpResponseCode(http.StatusInternalServerError).Responds()
					c.Abort()
					return
				}
				if !apiV1Util.Compare(sigGen.Signature, sigFromClient.Signature) {
					apiV1Util.ErrorPack(c).WithMessage("invalid signature").WithHttpResponseCode(http.StatusUnauthorized).Responds()
					c.Abort()
					return
				}
				c.Next()
			}
		} else {
			apiV1Util.ErrorPack(c).WithMessage("Signature extracting failed, maybe you should use signature middleware first.").WithHttpResponseCode(http.StatusForbidden).Responds()
			c.Abort()
		}
	}
}
