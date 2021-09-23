package v1

import (
	"github.com/gin-gonic/gin"
	"gol-c/database"
	"gol-c/model"
	"net/http"
)

type LoginJsonForm struct {
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
	loginJsonForm := &LoginJsonForm{}
	GetJsonForm(c, loginJsonForm)
	user := &model.User{}
	database.GetByField(&model.User{Username: loginJsonForm.Username}, user, []string{})
	pass := checkEncrypt(user.Password, loginJsonForm.Password)
	if pass {
		if user.EmailVerified {
			session := model.CreateSession(user.ID)
			err := database.Create(c, session, "session", ErrorMessageStatus)
			if err != nil {
				return
			}
			SuccessDataMessage(c, gin.H{"session_body": session.SessionBody}, "Create session succeeded!")
		} else {
			ErrorMessageStatus(c, "Create session failed: email not verified!", http.StatusUnauthorized)
		}
	} else {
		ErrorMessageStatus(c, "Create session failed: wrong username or password!", http.StatusUnauthorized)
	}
}
