package util

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetJsonForm(c *gin.Context, form interface{}) {
	err := c.BindJSON(form)

	if err != nil {
		log.Printf("JSON format not allowed!")
		ErrorPack(c).WithMessage("JSON format not allowed!").WithHttpResponseCode(http.StatusBadRequest).Responds()
		return
	}
}
