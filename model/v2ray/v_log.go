package v2ray

type LogLevel string

const (
	LogLevelDebug   LogLevel = "debug"
	LogLevelInfo    LogLevel = "info"
	LogLevelWarning LogLevel = "warning"
	LogLevelError   LogLevel = "error"
	LogLevelNone    LogLevel = "none"
)

// ConfLog https://www.v2ray.com/chapter_02/01_overview.html#logobject
type ConfLog struct {
	Access   string   `json:"access,omitempty"`
	Error    string   `json:"error,omitempty"`
	Loglevel LogLevel `json:"loglevel"`
}

func (l *ConfLog) AsDefault() *ConfLog {
	l.Loglevel = LogLevelInfo
	return l
}
