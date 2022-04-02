package model

import (
	"encoding/json"
	"ockham-api/model/v2ray"
)

type V2RayConfig struct {
	Stats     struct{}                   `json:"stats"` // 流量统计，存在即启用 https://www.v2ray.com/chapter_02/stats.html
	Api       *v2ray.ConfApi             `json:"api"`
	Log       *v2ray.ConfLog             `json:"log"`
	Policy    *v2ray.ConfPolicy          `json:"policy"`
	Inbounds  []*v2ray.ConfInboundsItem  `json:"inbounds"`
	Outbounds []*v2ray.ConfOutboundsItem `json:"outbounds"`
	Routing   *v2ray.ConfRouting         `json:"routing"`
}

func (vc *V2RayConfig) AsJSON() string {
	jsonData, _ := json.Marshal(vc)
	return string(jsonData)
}

func GenConfig(inboundPort, apiPort int, websocketPath string) V2RayConfig {
	return V2RayConfig{
		Stats:  struct{}{},
		Api:    (&v2ray.ConfApi{}).AsDefault(),
		Log:    (&v2ray.ConfLog{}).AsDefault(),
		Policy: (&v2ray.ConfPolicy{}).AsDefault(),
		Inbounds: []*v2ray.ConfInboundsItem{
			(&v2ray.ConfInboundsItem{}).AsInboundVmess(inboundPort, websocketPath),
			(&v2ray.ConfInboundsItem{}).AsDokodemo(apiPort),
		},
		Outbounds: []*v2ray.ConfOutboundsItem{
			(&v2ray.ConfOutboundsItem{}).AsFreedom(),
		},
		Routing: (&v2ray.ConfRouting{}).AsDefault(),
	}
}
