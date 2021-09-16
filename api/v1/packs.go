package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type meta struct {
	success bool
	message string
}

type Returner struct {
	meta meta
	data interface{}
}

func Success(c *gin.Context) {
	SuccessData(c, nil)
}

func SuccessData(c *gin.Context, data interface{}) {
	SuccessDataMessage(c, data, "OK!")
}

func SuccessDataMessage(c *gin.Context, data interface{}, message string) {
	SuccessDataMessageStatus(c, data, message, http.StatusOK)
}

func SuccessDataMessageStatus(c *gin.Context, data interface{}, message string, status int) {
	c.JSON(status, Returner{
		meta: meta{
			success: true,
			message: message,
		},
		data: data,
	})
}

func ErrorMessageStatus(c *gin.Context, message string, status int) {
	c.JSON(status, Returner{
		meta: meta{
			success: true,
			message: message,
		},
		data: nil,
	})
}
