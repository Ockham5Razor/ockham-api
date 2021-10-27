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

func HasAllAuthorities(neededAuthorities ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 设置为 SET
		//neededAuthoritySet := make(map[string]bool)
		//for i := range neededAuthorities {
		//	neededAuthoritySet[string.ToLower(neededAuthorities[i])] = true
		//}
		token, tokenExists := c.Get("token")
		if tokenExists {
			jwtClaims, _ := util.ParseToken(token.(string))
			username := jwtClaims.Username
			user := &model.User{}
			database.DBConn.Debug().Preload("Roles").First(user, &model.User{Username: username})
			hasAuthoritySet := make(map[string]bool)
			roles := user.Roles
			for i := range roles {
				fmt.Println(roles[i].RoleName)
				hasAuthoritySet[strings.ToLower(roles[i].RoleName)] = true
			}
			for i := range neededAuthorities {
				_, hasOneNeededAuthority := hasAuthoritySet[strings.ToLower(neededAuthorities[i])]
				if !hasOneNeededAuthority {
					apiV1Util.ErrorMessageStatus(c, "User not authorized as "+neededAuthorities[i], http.StatusForbidden)
					c.Abort()
				}
			}
			c.Next()
		} else {
			apiV1Util.ErrorMessageStatus(c, "Token not exists.", http.StatusBadRequest)
			c.Abort()
		}
	}
}
