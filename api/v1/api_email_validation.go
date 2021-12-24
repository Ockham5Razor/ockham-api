package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ockham-api/api/v1/util"
	"ockham-api/database"
	"ockham-api/model"
	"time"
)

type ValidateEmailJsonForm struct {
	ValidatorKey  string
	ValidatorCode string
}

// ValidateEmail
// @Summary ValidateEmail
// @SubscriptionDescription Validate a user's email
// @Tags auth
// @Success 200 {object} util.Pack
// @Failure 410 {object} util.Pack
// @Param param body ValidateEmailJsonForm true "Email validation from"
// @Router /v1/auth/email-validations/any/validating [PUT]
func ValidateEmail(c *gin.Context) {
	validateEmailJsonForm := &ValidateEmailJsonForm{}
	util.FillJsonForm(c, validateEmailJsonForm)

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
			util.SuccessPack(c).WithMessage("Email validated!").Responds()
		} else {
			util.ErrorPack(c).WithMessage("Email validating failed: validation expired.").WithHttpResponseCode(http.StatusGone).Responds()
		}
	} else {
		util.ErrorPack(c).WithMessage("Email validating failed: wrong validation key or code").WithHttpResponseCode(http.StatusGone).Responds()
	}
}
