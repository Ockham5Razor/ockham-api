package v1

import (
	"github.com/gin-gonic/gin"
	"gol-c/api/v1/util"
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
// @Tags auth
// @Success 200 {string} string    "ok"
// @Param param body LoginJsonForm true "CreateSession from"
// @Router /v1/auth/sessions [POST]
func CreateSession(c *gin.Context) {
	loginJsonForm := &LoginJsonForm{}
	util.GetJsonForm(c, loginJsonForm)
	user := &model.User{}
	database.GetByField(&model.User{Username: loginJsonForm.Username}, user, []string{})
	pass := util.CheckEncrypt(user.Password, loginJsonForm.Password)
	if pass {
		if user.EmailVerified {
			session := model.CreateSession(user.ID)
			err := database.Create(c, session, "session", util.ErrorMessageStatus)
			if err != nil {
				return
			}
			util.SuccessDataMessage(c, gin.H{"session_body": session.SessionBody}, "Create session succeeded!")
		} else {
			util.ErrorMessageStatus(c, "Create session failed: email not verified!", http.StatusUnauthorized)
		}
	} else {
		util.ErrorMessageStatus(c, "Create session failed: wrong username or password!", http.StatusUnauthorized)
	}
}
