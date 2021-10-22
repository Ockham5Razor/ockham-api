package v1

import (
	"github.com/gin-gonic/gin"
	"gol-c/api/v1/util"
	"gol-c/database"
	"gol-c/model"
	"net/http"
	"time"
)

type ValidateEmailJsonForm struct {
	ValidatorKey  string
	ValidatorCode string
}

// ValidateEmail
// @Summary ValidateEmail
// @Description Validate a user's email
// @Tags auth
// @Success 200 {string} string    "ok"
// @Param param body ValidateEmailJsonForm true "Email validation from"
// @Router /v1/auth/email-validations/any:validate [PUT]
func ValidateEmail(c *gin.Context) {
	validateEmailJsonForm := &ValidateEmailJsonForm{}
	util.GetJsonForm(c, validateEmailJsonForm)

	emailValidation := &model.EmailValidation{}
	database.GetByField(&model.EmailValidation{ValidationKey: validateEmailJsonForm.ValidatorKey}, &emailValidation, []string{"User"})

	if emailValidation.ValidationCode == validateEmailJsonForm.ValidatorCode {
		now := time.Now()
		if now.Before(emailValidation.ExpireAt) {
			err := database.Updates(c, emailValidation.User, &model.User{EmailVerified: true}, "user", util.ErrorMessageStatus)
			if err != nil {
				return
			}
			err = database.Delete(c, &model.EmailValidation{}, emailValidation.ID, "email validation", util.ErrorMessageStatus)
			if err != nil {
				return
			}
			util.SuccessMessage(c, "Email validated!")
		} else {
			util.ErrorMessageStatus(c, "Email validating failed: validation expired.", http.StatusGone)
		}
	} else {
		util.ErrorMessageStatus(c, "Email validating failed: wrong validation key or code", http.StatusGone)
	}
}
