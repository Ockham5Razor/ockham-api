package config

var (
	PortalListen                       = "0.0.0.0:8080"
	DbHost                             string
	DbPort                             int
	DbSchema                           string
	DbUser                             string
	DbPass                             string
	DbCharset                          string
	EmailValidationExpireDuration      = "30m"
	AuthSessionExpireSeconds           = 7200
	AuthSessionMaximumRenewalTimes     = 3
	JwtIssuer                          = "ockham-api"
	JwtSecret                          string
	JwtExpireSeconds                   = 1800
	SignatureTimestampToleranceSeconds = 60
	SignatureBodyDigestTruncateBytes   = 128
)
