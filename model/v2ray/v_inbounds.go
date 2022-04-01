package v2ray

import "ockham-api/model/v2ray/inbound_protocol"

type InboundProtocol string

const (
	InboundProtocolVmess        InboundProtocol = "vmess"
	InboundProtocolDokodemoDoor InboundProtocol = "dokodemo-door"
)

type ConfInboundsItemSettings interface{}

type ConfInboundsItemStreamSettingsWsSettings struct {
	Path    string            `json:"path,omitempty"`
	Headers map[string]string `json:"headers,omitempty"`
}

type ConfInboundsItemStreamSettings struct {
	Network    string                                    `json:"network,omitempty"`
	WsSettings *ConfInboundsItemStreamSettingsWsSettings `json:"wsSettings,omitempty"`
}

type ConfInboundsItemSniffing struct {
	Enabled      bool     `json:"enabled,omitempty"`
	DestOverride []string `json:"destOverride,omitempty"`
}

type ConfInboundsItemAllocate struct {
	Strategy    string `json:"strategy,omitempty"`
	Refresh     int    `json:"refresh,omitempty"`
	Concurrency int    `json:"concurrency,omitempty"`
}

// ConfInboundsItem https://www.v2ray.com/chapter_02/01_overview.html#inboundobject
type ConfInboundsItem struct {
	Protocol       InboundProtocol                 `json:"protocol"`
	Listen         string                          `json:"listen"`
	Port           int                             `json:"port"`
	Tag            string                          `json:"tag"`
	Settings       ConfInboundsItemSettings        `json:"settings"`
	StreamSettings *ConfInboundsItemStreamSettings `json:"streamSettings,omitempty"`
	Sniffing       *ConfInboundsItemSniffing       `json:"sniffing,omitempty"`
	Allocate       *ConfInboundsItemAllocate       `json:"allocate,omitempty"`
}

func (i *ConfInboundsItem) AsDokodemo() *ConfInboundsItem {
	i.Protocol = InboundProtocolDokodemoDoor
	i.Listen = "0.0.0.0"
	i.Port = 8080
	i.Tag = "api"
	i.Settings = inbound_protocol.ConfDokodemoDoorSettings{
		Address: "0.0.0.0",
		Port:    8080,
	}
	return i
}

func (i *ConfInboundsItem) AsInboundVmess() *ConfInboundsItem {
	i.Protocol = InboundProtocolVmess
	i.Listen = "0.0.0.0"
	i.Port = 10086
	i.Tag = "proxy"
	i.Settings = inbound_protocol.ConfVmessSettings{
		Clients:                   []inbound_protocol.ConfVmessSettingsClientsItem{},
		DisableInsecureEncryption: true, // 禁止不安全加密方式连接
	}
	i.StreamSettings = &ConfInboundsItemStreamSettings{
		Network: "ws",
		WsSettings: &ConfInboundsItemStreamSettingsWsSettings{
			Path: "/access-may/",
		},
	}
	return i
}
