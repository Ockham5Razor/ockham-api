package v2ray

type ConfPolicySystem struct {
	StatsInboundUplink   bool `json:"statsInboundUplink"`
	StatsInboundDownlink bool `json:"statsInboundDownlink"`
}

type ConfPolicyLevelsValue struct {
	Handshake         int  `json:"handshake,omitempty"`
	ConnIdle          int  `json:"connIdle,omitempty"`
	UplinkOnly        int  `json:"uplinkOnly,omitempty"`
	DownlinkOnly      int  `json:"downlinkOnly,omitempty"`
	StatsUserUplink   bool `json:"statsUserUplink,omitempty"`
	StatsUserDownlink bool `json:"StatsUserDownlink,omitempty"`
	BufferSize        int  `json:"bufferSize,omitempty"`
}

// ConfPolicy https://www.v2ray.com/chapter_02/policy.html
type ConfPolicy struct {
	System ConfPolicySystem                 `json:"system"`
	Levels map[string]ConfPolicyLevelsValue `json:"levels"`
}

func (p *ConfPolicy) AsDefault() *ConfPolicy {
	p.System = ConfPolicySystem{
		StatsInboundUplink:   true,
		StatsInboundDownlink: true,
	}
	p.Levels = map[string]ConfPolicyLevelsValue{
		"5": {
			StatsUserUplink:   true,
			StatsUserDownlink: true,
		},
	}
	return p
}
