package inbound_protocol

type domainStrategy string

// ConfDokodemoDoorSettings https://www.v2ray.com/chapter_02/protocols/dokodemo.html
type ConfDokodemoDoorSettings struct {
	Address        string `json:"address,omitempty"`
	Port           int    `json:"port,omitempty"`
	Network        string `json:"network,omitempty"`
	Timeout        int    `json:"timeout,omitempty"`
	FollowRedirect bool   `json:"FollowRedirect,omitempty"`
	UserLevel      int    `json:"userLevel,omitempty"`
}
