package v2ray

const (
	ApiTagDefault = "api"
)

var (
	ApiServicesOptions = []string{"HandlerService", "StatsService", "LoggerService"}
	ApiServicesDefault = []string{"HandlerService", "StatsService"}
)

// ConfApi https://www.v2ray.com/chapter_02/api.html
type ConfApi struct {
	Tag      string   `json:"tag"`
	Services []string `json:"services"`
}

func (a *ConfApi) AsDefault() *ConfApi {
	a.Tag = ApiTagDefault
	a.Services = ApiServicesDefault
	return a
}
