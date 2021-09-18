package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gol-c/database"
	"gol-c/model"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type RegisterJsonForm struct {
	FormJson
	Username string
	Password string
	Email    string
}

type LoginJsonForm struct {
	FormJson
	Username string
	Password string
}

// CreateUser
// @Summary CreateUser
// @Description Register to create a user
// @Success 200 {string} string    "ok"
// @Param param body RegisterJsonForm true "CreateUser from"
// @Router /v1/auth/users [POST]
func CreateUser(c *gin.Context) {
	registerJsonForm := RegisterJsonForm{}
	registerJsonForm.GetJsonForm(c)

	user := &model.User{
		Username: registerJsonForm.Username,
		Password: encrypt(registerJsonForm.Password),
		Email:    registerJsonForm.Email,
	}

	err := database.Create(c, user, "user", ErrorMessageStatus)
	if err != nil {
		return
	}

	SuccessDataMessageStatus(c, nil, "OK!", http.StatusCreated)
}

// CreateSession
// @Summary CreateSession
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

func checkEncrypt(hashed string, toCheckRawString string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(toCheckRawString))
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func encrypt(rawString string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(rawString), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	return string(hash)
}
