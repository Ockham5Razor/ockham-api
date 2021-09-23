package v1

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetJsonForm(c *gin.Context, form interface{}) {
	err := c.BindJSON(form)

	if err != nil {
		log.Printf("JSON format not allowed!")
		ErrorMessageStatus(c, "JSON format not allowed!", http.StatusBadRequest)
		return
	}
}
