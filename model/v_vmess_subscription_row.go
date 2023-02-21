package model

import (
	"encoding/json"
	"fmt"
	"ockham-api/util"
)

type VmessSubscriptionRow struct {
	Port uint   `json:"port"`
	Aid  uint   `json:"aid"`
	ID   string `json:"id"`
	Ps   string `json:"ps"`
	Add  string `json:"add"`
	Tls  string `json:"tls"`
	V    string `json:"v"`
	Net  string `json:"net"`
	Host string `json:"host"`
	Path string `json:"path"`
	Type string `json:"type"`
}

func (r *VmessSubscriptionRow) AsLink() string {
	rowJSONBytes, _ := json.Marshal(r)
	base64 := util.Base64EncodeBytes(rowJSONBytes)
	return fmt.Sprintf("vmess://%v", base64)
}
