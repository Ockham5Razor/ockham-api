package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Meta struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type Pack struct {
	context          *gin.Context
	httpResponseCode int
	Meta             Meta        `json:"meta"`
	Body             interface{} `json:"body"`
}

func SuccessPack(c *gin.Context) *Pack {
	return &Pack{
		context:          c,
		httpResponseCode: 200,
		Meta: Meta{
			Success: true,
			Message: "OK!",
		},
		Body: nil,
	}
}

func ErrorPack(c *gin.Context) *Pack {
	return &Pack{
		context:          c,
		httpResponseCode: 400,
		Meta: Meta{
			Success: false,
			Message: "ERROR!",
		},
		Body: nil,
	}
}

func (p *Pack) WithMessage(message string, a ...interface{}) *Pack {
	p.Meta.Message = fmt.Sprintf(message, a)
	return p
}

func (p *Pack) WithData(data interface{}) *Pack {
	p.Body = data
	return p
}

func (p *Pack) WithHttpResponseCode(responseCode int) *Pack {
	p.httpResponseCode = responseCode
	return p
}

func (p *Pack) Responds() {
	p.context.JSON(p.httpResponseCode, p)
}

func ErrorMessageStatus(c *gin.Context, message string, status int) {
	ErrorPack(c).WithMessage(message).WithHttpResponseCode(status).Responds()
}
