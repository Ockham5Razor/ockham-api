package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ockham-api/api/v1/util"
	"strings"
)

func Token() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenValueFromHeader := c.Request.Header.Get(util.HeaderToken)
		if tokenValueFromHeader == "" {
			util.ErrorPack(c).WithMessage("HTTP header `" + util.HeaderToken + "` is required for this request.").WithHttpResponseCode(http.StatusBadRequest).Responds()
			c.Abort()
			return
		}
		parts := strings.SplitN(tokenValueFromHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			util.ErrorPack(c).WithMessage("HTTP header `" + util.HeaderToken + "` is in illegal format.").WithHttpResponseCode(http.StatusBadRequest).Responds()
			c.Abort()
			return
		}
		token := parts[1]
		c.Set(util.ContextBearerValue, token)
		c.Next()
	}
}

func Signature() gin.HandlerFunc {
	return func(c *gin.Context) {
		signatureValueFromHeader := c.Request.Header.Get(util.HeaderToken)
		if signatureValueFromHeader == "" {
			util.ErrorPack(c).WithMessage("HTTP header `" + util.HeaderToken + "` is required for this request.").WithHttpResponseCode(http.StatusBadRequest).Responds()
			c.Abort()
			return
		}
		timestampValueFromHeader := c.Request.Header.Get(util.HeaderTimestamp)
		if timestampValueFromHeader == "" {
			util.ErrorPack(c).WithMessage("HTTP header `" + util.HeaderTimestamp + "` is required for this request.").WithHttpResponseCode(http.StatusBadRequest).Responds()
			c.Abort()
			return
		}
		parts := strings.SplitN(signatureValueFromHeader, " ", 2)
		if !(parts[0] == "Signature") {
			util.ErrorPack(c).WithMessage("HTTP header `" + util.HeaderToken + "` is in illegal format.").WithHttpResponseCode(http.StatusBadRequest).Responds()
			c.Abort()
			return
		}

		c.Set(util.ContextSignatureValue, signatureValueFromHeader)
		c.Next()
	}
}
