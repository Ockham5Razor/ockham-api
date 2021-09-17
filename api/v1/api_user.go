package v1

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
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
		return
	}

	user := &model.User{
		Username: registerJsonForm.Username,
		Password: encrypt(registerJsonForm.Password),
		Email:    registerJsonForm.Email,
	}

	if dbc := database.DBConn.Create(user); dbc.Error != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(dbc.Error, &mysqlErr) && mysqlErr.Number == 1062 {
			ErrorMessageStatus(c, "Create user failed: user already exists.", http.StatusBadRequest)
			return
		}
		ErrorMessageStatus(c, "Create user failed: unknown.", http.StatusBadRequest)
		return
	}

	Success(c)
}

func encrypt(rawPassword string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	return string(hash)
}
