package model

import (
	"errors"
	"gorm.io/gorm"
	"ockham-api/database"
	"ockham-api/util"
	"testing"
	"time"
)

func TestInsertBilling(t *testing.T) {
	// init database
	_ = util.InitDatabase()

	sIDList := IDList{1, 12, 333}
	tIDList := IDList{21, 32, 555}

	billing := Billing{
		BillingTitle:            "test",
		BillingDescription:      "test",
		BillingTotal:            100.00,
		BillingDate:             time.Now(),
		PaymentSettled:          false,
		PaymentStartDate:        time.Now(),
		PaymentDueDate:          time.Now(),
		SubscribingServicePlans: &sIDList,
		SubscribingTrafficPlans: &tIDList,
	}

	// rollback when testing succeeded.
	database.DBConn.Transaction(func(tx *gorm.DB) error {
		tx.Create(&billing)
		return errors.New("common testing succeeded and rollback")
	})
}
