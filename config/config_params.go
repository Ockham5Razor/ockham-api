package config

var (
	PortalListen                           = "0.0.0.0:8080"
	DbHost                                 string
	DbPort                                 int
	DbSchema                               string
	DbUser                                 string
	DbPass                                 string
	DbCharset                              string
	EmailUser                              string
	EmailPass                              string
	EmailHost                              string
	EmailPort                              int
	EmailSign                              string
	EmailValidationExpireDuration          string
	AuthSessionExpireSeconds               int
	AuthSessionMaximumRenewalTimes         int
	AuthJwtIssuer                          string
	AuthJwtSecret                          string
	AuthJwtExpireSeconds                   int
	AuthSignatureTimestampToleranceSeconds = 60
	AuthSignatureBodyDigestTruncateBytes   = 128
)
