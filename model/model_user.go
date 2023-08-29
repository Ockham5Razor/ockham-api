package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"ockham-api/database"
)

type User struct {
	gorm.Model    `json:"-"`
	Username      string  `gorm:"type:VARCHAR(24);uniqueIndex" json:"username"`
	Password      string  `gorm:"type:VARCHAR(128)" json:"-"`
	Email         string  `gorm:"type:VARCHAR(128)" json:"email"`
	EmailVerified bool    `json:"email_verified"`
	Roles         []*Role `gorm:"many2many:user_role;" json:"-"`
}

type Role struct {
	gorm.Model `json:"-"`
	RoleName   string `gorm:"type:VARCHAR(24)" json:"role_name"`
}

func GetUser(userID uint64) *User {
	user := &User{}
	database.DBConn.Preload("Roles").First(user, userID)
	return user
}

func (user *User) RemoveRole(roleId uint) {
	targetIndex := -1
	for i, role := range user.Roles {
		if role.ID == roleId {
			fmt.Println(i)
			targetIndex = i
		}
	}
	fmt.Println(user.Roles)
	if targetIndex != -1 {
		user.Roles = append(user.Roles[:targetIndex], user.Roles[targetIndex+1:]...)
	}
	fmt.Println(user.Roles)
	_ = database.DBConn.Model(user).Association("Roles").Replace(user.Roles)
}

func (user *User) GetJSON() gin.H {
	//database.GetByField()
	return gin.H{
		"id":       user.ID,
		"username": user.Username,
		"nickname": user.Username,
		"roles":    user.Roles,
	}
}
