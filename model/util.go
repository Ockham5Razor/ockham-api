package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type IDList []uint

// Scan scan value into Jsonb, implements sql.Scanner interface
func (idList *IDList) Scan(value interface{}) error {
	val, ok := value.(string)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal string value:", value))
	}
	err := json.Unmarshal([]byte(val), &idList)
	return err
}

// Value return json value, implement driver.Valuer interface
func (idList *IDList) Value() (driver.Value, error) {
	valBytes, err := json.Marshal(idList)
	return string(valBytes), err
}
