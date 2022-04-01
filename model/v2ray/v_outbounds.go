package v2ray

type OutboundProtocol string

const (
	OutboundProtocolFreedom OutboundProtocol = "freedom"
)

type ConfOutboundsItem struct {
	Protocol OutboundProtocol `json:"protocol"`
	Settings struct{}         `json:"settings"`
}

func (o *ConfOutboundsItem) AsFreedom() *ConfOutboundsItem {
	o.Protocol = OutboundProtocolFreedom
	return o
}
