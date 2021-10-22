package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Success(c *gin.Context) {
	SuccessData(c, nil)
}

func SuccessData(c *gin.Context, data interface{}) {
	SuccessDataMessage(c, data, "OK!")
}

func SuccessMessage(c *gin.Context, message string) {
	SuccessDataMessageStatus(c, nil, message, http.StatusOK)
}

func SuccessDataMessage(c *gin.Context, data interface{}, message string) {
	SuccessDataMessageStatus(c, data, message, http.StatusOK)
}

func SuccessDataMessageStatus(c *gin.Context, data interface{}, message string, status int) {
	c.JSON(status, gin.H{
		"meta": gin.H{
			"success": false,
			"message": message,
		},
		"data": data,
	})
}

func ErrorMessageStatus(c *gin.Context, message string, status int) {
	c.JSON(status, gin.H{
		"meta": gin.H{
			"success": false,
			"message": message,
		},
		"data": nil,
	})
}
