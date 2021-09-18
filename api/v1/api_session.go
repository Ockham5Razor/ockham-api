package v1

import (
	"github.com/gin-gonic/gin"
	"gol-c/database"
	"gol-c/model"
	"net/http"
)

type LoginJsonForm struct {
	FormJson
	Username string
	Password string
}

// CreateSession
// @Summary Login
// @Description Login as a user
// @Success 200 {string} string    "ok"
// @Param param body LoginJsonForm true "CreateSession from"
// @Router /v1/auth/sessions [POST]
func CreateSession(c *gin.Context) {
	loginJsonForm := LoginJsonForm{}
	loginJsonForm.GetJsonForm(c)
	user := &model.User{}
	database.GetByField(&model.User{Username: loginJsonForm.Username}, user)
	pass := checkEncrypt(user.Password, loginJsonForm.Password)
	if pass {
		session := model.CreateSession(user.ID)
		err := database.Create(c, session, "session", ErrorMessageStatus)
		if err != nil {
			return
		}
		SuccessDataMessage(c, gin.H{"session_body": session.SessionBody}, "CreateSession Succeeded!")
	} else {
		ErrorMessageStatus(c, "CreateSession Failed!", http.StatusUnauthorized)
	}
}
