package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ockham-api/api/v1/util"
	"ockham-api/database"
	"ockham-api/model"
	"strconv"
)

type RegisterJsonForm struct {
	Username string
	Password string
	Email    string
}

// CreateUser
// @Summary Register
// @SubscriptionDescription Register to create a user
// @Tags auth
// @Success 201 {object} util.Pack
// @Failure 409,500 {object} util.Pack
// @Param param body RegisterJsonForm true "CreateUser from"
// @Router /v1/auth/users [POST]
func CreateUser(c *gin.Context) {
	registerJsonForm := &RegisterJsonForm{}
	util.FillJsonForm(c, registerJsonForm)

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

// GetMe
// @Summary Get Current User
// @SubscriptionDescription Get current user
// @Tags auth
// @Security Bearer
// @Success 200 {object} util.Pack
// @Failure 409,500 {object} util.Pack
// @Router /v1/auth/users/me [GET]
func GetMe(c *gin.Context) {
	user, _ := c.Get("user")
	util.SuccessPack(c).WithHttpResponseCode(http.StatusOK).WithData(user).Responds()
}

type GrantRoleForm struct {
	RoleID uint
}

// GrantRole
// @Summary Grant Role
// @SubscriptionDescription Grant A Role to A User
// @Tags auth
// @Security Bearer
// @Success 201 {object} util.Pack
// @Failure 409,500 {object} util.Pack
// @Param param body GrantRoleForm true "GrantRoleForm from"
// @Param user_id path int true "user id"
// @Router /v1/auth/users/{user_id}/roles [POST]
func GrantRole(c *gin.Context) {
	grantRoleForm := &GrantRoleForm{}
	util.FillJsonForm(c, grantRoleForm)

	userIdStr := c.Param("user_id")
	userIdU64, _ := strconv.ParseUint(userIdStr, 10, 32)

	targetUser := &model.User{}
	database.Get(uint(userIdU64), targetUser)

	targetRole := &model.Role{}
	database.Get(grantRoleForm.RoleID, targetRole)

	targetUser.Roles = append(targetUser.Roles, targetRole)
	_ = database.Update(c, targetUser, "User Role", util.ErrorMessageStatus)

	util.SuccessPack(c).WithHttpResponseCode(http.StatusCreated).Responds()
}

// RevokeRole
// @Summary Revoke Role
// @SubscriptionDescription Revoke A Role from A User
// @Tags auth
// @Security Bearer
// @Success 201 {object} util.Pack
// @Failure 409,500 {object} util.Pack
// @Param user_id path int true "user id"
// @Param role_id path int true "role id"
// @Router /v1/auth/users/{user_id}/roles/{role_id} [DELETE]
func RevokeRole(c *gin.Context) {
	userIdStr := c.Param("user_id")
	userIdU64, _ := strconv.ParseUint(userIdStr, 10, 32)

	roleIdStr := c.Param("role_id")
	roleIdU64, _ := strconv.ParseUint(roleIdStr, 10, 32)

	targetUser := model.GetUser(userIdU64)
	targetUser.RemoveRole(uint(roleIdU64))

	util.SuccessPack(c).WithHttpResponseCode(http.StatusCreated).Responds()
}
