package outbound_protocol

type domainStrategy string

const (
	AsIs    domainStrategy = "AsIs"
	UseIP   domainStrategy = "UseIP"
	UseIPv4 domainStrategy = "UseIPv4"
	UseIPv6 domainStrategy = "UseIPv6"
)

// ConfFreedomSettings https://www.v2ray.com/chapter_02/protocols/freedom.html
type ConfFreedomSettings struct {
	DomainStrategy *domainStrategy `json:"domainStrategy,omitempty"`
	Redirect       string          `json:"redirect,omitempty"`
	UserLevel      int             `json:"userLevel,omitempty"`
}
