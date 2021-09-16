package v1

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type RegisterJsonForm struct {
	Username    string
	RawPassword string
	Email       string
}

// Register
// @Summary Register
// @Description Register to create a user
// @Success 200 {string} string    "ok"
// @Param param body RegisterJsonForm true "Register from"
// @Router /v1/auth/register [POST]
func Register(c *gin.Context) {
	registerJsonForm := RegisterJsonForm{}
	err := c.BindJSON(&registerJsonForm)

	if err != nil {
		log.Printf("JSON format not allowed!")
		c.JSON(http.StatusBadRequest, gin.H{})
	}

	Success(c)
}
