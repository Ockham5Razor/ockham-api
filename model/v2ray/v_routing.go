package v2ray

type ConfRoutingRulesItem struct {
	InboundTag  []string `json:"inboundTag"`
	OutboundTag string   `json:"outboundTag"`
	Type        string   `json:"type"`
}

type ConfRoutingSetting struct {
	Rules []*ConfRoutingRulesItem `json:"rules"`
}

// ConfRouting https://www.v2ray.com/chapter_02/03_routing.html
// 此版本 model 为测试可用的版本，与文档中的新版本可能不一致！
type ConfRouting struct {
	Strategy string              `json:"strategy"`
	Setting  *ConfRoutingSetting `json:"setting"`
}

func (r *ConfRouting) AsDefault() *ConfRouting {
	r.Strategy = "rules"
	r.Setting = &ConfRoutingSetting{
		Rules: []*ConfRoutingRulesItem{
			{
				InboundTag:  []string{"api"},
				OutboundTag: "api",
				Type:        "field",
			},
		},
	}
	return r
}
