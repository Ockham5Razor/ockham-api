package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ockham-api/api/v1/util"
	"ockham-api/database"
	"ockham-api/model"
)

type RechargeCodeGeneratingJsonForm struct {
	PackageName    string
	BatchCount     int
	RechargeAmount float32
}

// GenerateRechargeCodesInBatches
// @Summary Generate Recharge Codes in Batches
// @SubscriptionDescription Generate Recharge Codes in Batches
// @Tags recharge_code
// @Security Bearer
// @Success 201 {object} util.Pack
// @Failure 409,500 {object} util.Pack
// @Param param body RechargeCodeGeneratingJsonForm true "Generating recharge code from"
// @Router /v1/recharge-codes/none/batch-generations [POST]
func GenerateRechargeCodesInBatches(c *gin.Context) {
	genForm := &RechargeCodeGeneratingJsonForm{}
	util.FillJsonForm(c, genForm)
	var rechargeCodes []*model.RechargeCode
	for i := 0; i < genForm.BatchCount; i++ {
		rechargeCodes = append(rechargeCodes, model.GenRechargeCode(genForm.PackageName, genForm.RechargeAmount))
	}
	err := database.CreateInBatches(c, rechargeCodes, len(rechargeCodes), "RechargeCode", util.ErrorMessageStatus)
	if err != nil {
		return
	}
	util.SuccessPack(c).WithMessage("Generate recharge codes in batches succeeded!").WithHttpResponseCode(http.StatusCreated).Responds()
}
