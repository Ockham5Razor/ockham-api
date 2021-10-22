package v1

import (
	"github.com/gin-gonic/gin"
	"gol-c/database"
	"gol-c/model"
	"net/http"
)

type RegisterJsonForm struct {
	Username string
	Password string
	Email    string
}

// CreateUser
// @Summary Register
// @Description Register to create a user
// @Tags auth
// @Success 200 {string} string    "ok"
// @Param param body RegisterJsonForm true "CreateUser from"
// @Router /v1/auth/users [POST]
func CreateUser(c *gin.Context) {
	registerJsonForm := &RegisterJsonForm{}
	GetJsonForm(c, registerJsonForm)

	user := &model.User{
		Username:      registerJsonForm.Username,
		Password:      encrypt(registerJsonForm.Password),
		Email:         registerJsonForm.Email,
		EmailVerified: false,
	}

	err := database.Create(c, user, "user", ErrorMessageStatus)
	if err != nil {
		return
	}

	emailVerification := model.NewEmailValidation(user)
	err = database.Create(c, emailVerification, "email verification", ErrorMessageStatus)
	if err != nil {
		return
	}

	SuccessDataMessageStatus(c, nil, "OK!", http.StatusCreated)
}
