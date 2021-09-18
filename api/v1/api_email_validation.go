package v1

import "github.com/gin-gonic/gin"

type ValidateEmailJsonForm struct {
	FormJson
	ValidatorKey  string
	ValidatorCode string
}

// ValidateEmail
// @Summary ValidateEmail
// @Description Register to create a user
// @Success 200 {string} string    "ok"
// @Param param body RegisterJsonForm true "CreateUser from"
// @Router /v1/auth/email-validation/validate [PUT]
func ValidateEmail(c *gin.Context) {
	validateEmailJsonForm := &ValidateEmailJsonForm{}
	validateEmailJsonForm.GetJsonForm(c)
}
