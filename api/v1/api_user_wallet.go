package v1

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"ockham-api/api/v1/util"
	"ockham-api/database"
	"ockham-api/model"
)

// GetWalletInfo
// @Summary Get wallet info
// @SubscriptionDescription Get wallet info
// @Tags wallet
// @Security Bearer
// @Success 200 {object} util.Pack
// @Failure 400 {object} util.Pack
// @Router /v1/users/me/wallet [GET]
func GetWalletInfo(c *gin.Context) {
	user, userExists := c.Get("user")
	if userExists {
		userObj := user.(*model.User)
		wallet := &model.UserWallet{}
		database.DBConn.FirstOrCreate(wallet, model.UserWallet{UserID: userObj.ID})
		util.SuccessPack(c).WithData(wallet).Responds()
	} else {
		util.ErrorPack(c).WithData("Token extracting failed, maybe you should use current user middleware first.").WithHttpResponseCode(http.StatusBadRequest).Responds()
		c.Abort()
	}
}

// GetRecordsOfWallet
// @Summary Get wallet records
// @SubscriptionDescription Get wallet records
// @Tags wallet
// @Security Bearer
// @Success 200 {object} util.Pack
// @Failure 400 {object} util.Pack
// @Router /v1/users/me/wallet/records [GET]
func GetRecordsOfWallet(c *gin.Context) {
	user, userExists := c.Get("user")
	if userExists {
		userObj := user.(*model.User)
		walletRecords := &[]model.UserWalletRecord{}
		database.DBConn.Order("created_at DESC").Find(walletRecords, model.UserWalletRecord{UserID: userObj.ID})
		util.SuccessPack(c).WithData(walletRecords).Responds()
	} else {
		util.ErrorPack(c).WithMessage("Token extracting failed, maybe you should use current user middleware first.").WithHttpResponseCode(http.StatusBadRequest).Responds()
		c.Abort()
	}
}

type RechargeForm struct {
	RechargeCode string
}

// RechargeWallet
// @Summary Recharge
// @SubscriptionDescription Register to create a user
// @Tags wallet
// @security Bearer
// @Success 201 {object} util.Pack
// @Failure 403,410 {object} util.Pack
// @Param param body RechargeForm true "Recharge form"
// @Router /v1/users/me/wallet/recharging [PUT]
func RechargeWallet(c *gin.Context) {
	user, userExists := c.Get("user")
	if userExists {
		database.DBConn.Begin()
		userObj := user.(*model.User)
		userID := userObj.ID

		// 获取前端传来的充值码
		rechargeForm := &RechargeForm{}
		util.FillJsonForm(c, rechargeForm)

		err := recharge(rechargeForm.RechargeCode, userID)
		if err == nil {
			util.SuccessPack(c).WithMessage("Recharging succeeded!").WithHttpResponseCode(http.StatusCreated).Responds()
		} else {
			util.ErrorPack(c).WithMessage(fmt.Sprintf("Recharging failed: %s.", err.Error())).WithHttpResponseCode(http.StatusGone).Responds()
		}
	} else {
		util.ErrorPack(c).WithMessage("Token extracting failed, maybe you should use current user middleware first.").WithHttpResponseCode(http.StatusBadRequest).Responds()
	}
}

func recharge(rechargeCodeString string, userID uint) error {
	err := database.DBConn.Transaction(func(tx *gorm.DB) error {

		// 获取用户的 wallet 对象
		wallet := &model.UserWallet{}
		database.DBConn.FirstOrCreate(wallet, &model.UserWallet{UserID: userID})

		// 获取充值码
		rechargeCode := &model.RechargeCode{}
		database.DBConn.First(rechargeCode, model.RechargeCode{RechargeCode: rechargeCodeString})

		if rechargeCode.Used {
			return errors.New("recharge code is used")
		} else {
			// 充值码标记为已使用
			rechargeCode.Used = true
			database.DBConn.Save(rechargeCode)

			// 余额增长
			wallet.Balance += rechargeCode.RechargeAmount
			database.DBConn.Save(wallet)

			// 流水记录
			database.DBConn.Create(&model.UserWalletRecord{
				Amount:       rechargeCode.RechargeAmount,
				Description:  "使用充值码充值",
				UserID:       userID,
				UserWalletID: wallet.ID,
			})

			return nil
		}
	})
	return err
}
