package v1

import (
	"github.com/gin-gonic/gin"
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
// @Success 200 {string} string    "ok"
// @Param param body ValidateEmailJsonForm true "Email validation from"
// @Router /v1/auth/email-validation/validate [PUT]
func ValidateEmail(c *gin.Context) {
	validateEmailJsonForm := &ValidateEmailJsonForm{}
	GetJsonForm(c, validateEmailJsonForm)

	emailValidation := &model.EmailValidation{}
	database.GetByField(&model.EmailValidation{ValidatorKey: validateEmailJsonForm.ValidatorKey}, &emailValidation, []string{"User"})

	if emailValidation.ValidatorCode == validateEmailJsonForm.ValidatorCode {
		now := time.Now()
		if now.Before(emailValidation.ExpireAt) {
			err := database.Updates(c, emailValidation.User, &model.User{EmailVerified: true}, "user", ErrorMessageStatus)
			if err != nil {
				return
			}
			err = database.Delete(c, &model.EmailValidation{}, emailValidation.ID, "email validation", ErrorMessageStatus)
			if err != nil {
				return
			}
			SuccessMessage(c, "Email validated!")
		} else {
			ErrorMessageStatus(c, "Email validating failed: validation expired.", http.StatusGone)
		}
	} else {
		ErrorMessageStatus(c, "Email validating failed: wrong validation key or code", http.StatusGone)
	}
}
