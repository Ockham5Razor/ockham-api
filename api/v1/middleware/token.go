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
			util.ErrorPack(c).WithMessage("HTTP header \"%v\" is required for this request.", util.HeaderToken).WithHttpResponseCode(http.StatusBadRequest).Responds()
			c.Abort()
			return
		}
		parts := strings.SplitN(tokenValueFromHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			util.ErrorPack(c).WithMessage("HTTP header \"%v\" is in illegal format.", util.HeaderToken).WithHttpResponseCode(http.StatusBadRequest).Responds()
			c.Abort()
			return
		}
		token := parts[1]
		c.Set(util.ContextBearerBody, token)
		c.Next()
	}
}

func Signature() gin.HandlerFunc {
	return func(c *gin.Context) {
		signatureValueFromHeader := c.Request.Header.Get(util.HeaderToken)
		if signatureValueFromHeader == "" {
			util.ErrorPack(c).WithMessage("HTTP header \"%v\" is required for this request.", util.HeaderToken).WithHttpResponseCode(http.StatusBadRequest).Responds()
			c.Abort()
			return
		}
		timestampValueFromHeader := c.Request.Header.Get(util.HeaderTimestamp)
		if timestampValueFromHeader == "" {
			util.ErrorPack(c).WithMessage("HTTP header \"%v\" is required for this request.", util.HeaderTimestamp).WithHttpResponseCode(http.StatusBadRequest).Responds()
			c.Abort()
			return
		}

		c.Set(util.ContextTimestampValue, timestampValueFromHeader)
		c.Set(util.ContextSignatureValue, signatureValueFromHeader)
		c.Next()
	}
}
