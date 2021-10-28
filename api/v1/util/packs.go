package util

import (
	"github.com/gin-gonic/gin"
)

type meta struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type Pack struct {
	context          *gin.Context
	httpResponseCode int
	Meta             meta        `json:"meta"`
	Data             interface{} `json:"data"`
}

func SuccessPack(c *gin.Context) *Pack {
	return &Pack{
		context:          c,
		httpResponseCode: 200,
		Meta: meta{
			Success: true,
			Message: "OK!",
		},
		Data: nil,
	}
}

func ErrorPack(c *gin.Context) *Pack {
	return &Pack{
		context:          c,
		httpResponseCode: 400,
		Meta: meta{
			Success: false,
			Message: "ERROR!",
		},
		Data: nil,
	}
}

func (p *Pack) WithMessage(message string) *Pack {
	p.Meta.Message = message
	return p
}

func (p *Pack) WithData(data interface{}) *Pack {
	p.Data = data
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
