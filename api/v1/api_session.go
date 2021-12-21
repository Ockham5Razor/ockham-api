package v1

import (
	"github.com/gin-gonic/gin"
	"gol-c/api/v1/util"
	"gol-c/database"
	"gol-c/model"
	s_util "gol-c/util"
	"net/http"
	"time"
)

type LoginJsonForm struct {
	Username string
	Password string
}

// CreateSession
// @Summary Login
// @SubscriptionDescription Login as a user
// @Tags auth
// @Success 201 {object} util.Pack
// @Failure 401,409,500 {object} util.Pack
// @Param param body LoginJsonForm true "Login json form"
// @Router /v1/auth/sessions [POST]
func CreateSession(c *gin.Context) {
	loginJsonForm := &LoginJsonForm{}
	util.FillJsonForm(c, loginJsonForm)
	user := &model.User{}
	database.GetByField(&model.User{Username: loginJsonForm.Username}, user, nil)
	pass := util.CheckEncrypt(user.Password, loginJsonForm.Password)
	if pass {
		if user.EmailVerified {
			session := model.CreateSession(user)
			err := database.Create(c, session, "session", util.ErrorMessageStatus)
			if err != nil {
				return
			}
			util.SuccessPack(c).WithData(gin.H{"renewal_code": session.SessionKey, "token_body": session.SessionToken}).WithHttpResponseCode(http.StatusCreated).WithMessage("Create session succeeded!").Responds()
		} else {
			util.ErrorPack(c).WithMessage("Create session failed: email not verified!").WithHttpResponseCode(http.StatusUnauthorized).Responds()
		}
	} else {
		util.ErrorPack(c).WithMessage("Create session failed: wrong username or password!").WithHttpResponseCode(http.StatusUnauthorized).Responds()
	}
}

type RenewSessionForm struct {
	RenewalKey string
}

// RenewSession
// @Summary Keep login status
// @SubscriptionDescription Keep login status as a user.
// @Tags auth
// @Success 201 {object} util.Pack
// @Failure 403,409,500 {object} util.Pack
// @Param param body RenewSessionForm true "session renewal form"
// @Router /v1/auth/sessions/any/renewing [PUT]
func RenewSession(c *gin.Context) {
	renewSessionForm := &RenewSessionForm{}
	util.FillJsonForm(c, renewSessionForm)
	session := &model.Session{}
	database.GetByField(&model.Session{SessionKey: renewSessionForm.RenewalKey}, session, nil)
	if time.Now().Before(session.ExpiredAt) {
		if session.RenewalStock > 0 || session.RenewalStock == -1 {
			body, _ := s_util.GenToken(session.User.Username, session.SessionKey)
			session.SessionToken = body
			session.RenewalStock -= 1
			err := database.Update(c, session, "session", util.ErrorMessageStatus)
			if err != nil {
				return
			}
			util.SuccessPack(c).WithData(gin.H{"renewal_code": session.SessionKey, "token_body": session.SessionToken}).WithMessage("Renew session succeeded!").WithHttpResponseCode(http.StatusCreated).Responds()
		} else {
			util.ErrorPack(c).WithMessage("Session renewing failed: reached maximum renewal times.").WithHttpResponseCode(http.StatusForbidden)
		}
	} else {
		util.ErrorPack(c).WithMessage("Session renewing failed: session expired.").WithHttpResponseCode(http.StatusForbidden)
	}
}
