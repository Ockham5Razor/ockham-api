package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gol-c/database"
	"gol-c/model"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

type RegisterJsonForm struct {
	Username string
	Password string
	Email    string
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
		ErrorMessageStatus(c, "JSON format not allowed!", http.StatusBadRequest)
	}

	user := &model.User{
		Username: registerJsonForm.Username,
		Password: encrypt(registerJsonForm.Password),
		Email:    registerJsonForm.Email,
	}

	database.DBConn.Save(user)

	Success(c)
}

func encrypt(rawPassword string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	return string(hash)
}
