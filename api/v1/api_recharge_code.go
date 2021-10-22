package v1

import (
	"github.com/gin-gonic/gin"
	"gol-c/database"
	"gol-c/model"
)

type RechargeCodeGeneratingJsonForm struct {
	PackageName    string
	BatchCount     int
	RechargeAmount float32
}

// GenerateRechargeCodesInBatches
// @Summary Generate Recharge Codes in Batches
// @Description Generate Recharge Codes in Batches
// @Tags recharge_code
// @Success 200 {string} string    "ok"
// @Param param body RechargeCodeGeneratingJsonForm true "Generating recharge code from"
// @Router /v1/recharge-codes/none/batch-generations [POST]
func GenerateRechargeCodesInBatches(c *gin.Context) {
	genForm := &RechargeCodeGeneratingJsonForm{}
	GetJsonForm(c, genForm)
	var rechargeCodes []*model.RechargeCode
	for i := 0; i < genForm.BatchCount; i++ {
		rechargeCodes = append(rechargeCodes, model.GenRechargeCode(genForm.PackageName, genForm.RechargeAmount))
	}
	err := database.CreateInBatches(c, rechargeCodes, len(rechargeCodes), "RechargeCode", ErrorMessageStatus)
	if err != nil {
		return
	}
	SuccessDataMessage(c, gin.H{}, "Generate recharge codes in batches succeeded!")
}
