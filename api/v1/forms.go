package v1

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type FormJson struct {
}

func (form *FormJson) GetJsonForm(c *gin.Context) {
	err := c.BindJSON(form)

	if err != nil {
		log.Printf("JSON format not allowed!")
		ErrorMessageStatus(c, "JSON format not allowed!", http.StatusBadRequest)
		return
	}
}
