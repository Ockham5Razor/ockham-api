package v1

import (
	"github.com/gin-gonic/gin"
	"gol-c/api/v1/util"
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
	util.GetJsonForm(c, registerJsonForm)

	user := &model.User{
		Username:      registerJsonForm.Username,
		Password:      util.Encrypt(registerJsonForm.Password),
		Email:         registerJsonForm.Email,
		EmailVerified: false,
	}

	err := database.Create(c, user, "user", util.ErrorMessageStatus)
	if err != nil {
		return
	}

	emailVerification := model.NewEmailValidation(user)
	err = database.Create(c, emailVerification, "email verification", util.ErrorMessageStatus)
	if err != nil {
		return
	}

	util.SuccessPack(c).WithHttpResponseCode(http.StatusCreated).Responds()
}
