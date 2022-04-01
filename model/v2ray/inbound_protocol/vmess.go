package inbound_protocol

type ConfVmessSettingsClientsItem struct {
	Id      string `json:"id"`
	Level   int    `json:"level"`
	AlterId int    `json:"alterId"`
	Email   string `json:"email"`
}

type ConfVmessSettingsDefault struct {
	Level   int `json:"level,omitempty"`
	AlterId int `json:"alterId,omitempty"`
}

type ConfVmessSettingsDetour struct {
	To string `json:"to,omitempty"`
}

// ConfVmessSettings https://www.v2ray.com/chapter_02/protocols/vmess.html#inboundconfigurationobject
type ConfVmessSettings struct {
	Clients                   []ConfVmessSettingsClientsItem `json:"clients"`
	Default                   *ConfVmessSettingsDefault      `json:"default,omitempty"`
	Detour                    *ConfVmessSettingsDetour       `json:"detour,omitempty"`
	DisableInsecureEncryption bool                           `json:"disableInsecureEncryption"`
}
