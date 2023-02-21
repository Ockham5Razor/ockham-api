package config

import (
	"github.com/spf13/viper"
)

var (
	DbHost    string
	DbPort    int
	DbSchema  string
	DbUser    string
	DbPass    string
	DbCharset string

	EmailUser string
	EmailPass string
	EmailHost string
	EmailPort int
	EmailSign string

	EmailValidationExpireDuration string

	AuthSessionExpireSeconds       int
	AuthSessionMaximumRenewalTimes int

	AuthJwtIssuer        string
	AuthJwtSecret        string
	AuthJwtExpireSeconds int

	AuthSignatureTimestampToleranceSeconds = 60
	AuthSignatureBodyDigestTruncateBytes   = 128
)

func FillParams() {
	DbHost = viper.GetString("db.host")
	DbPort = viper.GetInt("db.port")
	DbSchema = viper.GetString("db.schema")
	DbUser = viper.GetString("db.user")
	DbPass = viper.GetString("db.pass")
	DbCharset = viper.GetString("db.charset")

	EmailUser = viper.GetString("email.user")
	EmailPass = viper.GetString("email.pass")
	EmailHost = viper.GetString("email.host")
	EmailPort = viper.GetInt("email.port")
	EmailSign = viper.GetString("email.sign")

	EmailValidationExpireDuration = viper.GetString("email_validation.expire_duration")

	AuthSessionExpireSeconds = viper.GetInt("auth.session.expire_seconds")
	AuthSessionMaximumRenewalTimes = viper.GetInt("auth.session.maximum_renewal_times")

	AuthJwtIssuer = viper.GetString("auth.jwt.issuer")
	AuthJwtSecret = viper.GetString("auth.jwt.secret")
	AuthJwtExpireSeconds = viper.GetInt("auth.jwt.expire_seconds")

	AuthSignatureTimestampToleranceSeconds = viper.GetInt("auth.signature.timestamp_tolerance_seconds")
	AuthSignatureBodyDigestTruncateBytes = viper.GetInt("auth.signature.body_digest_truncate_bytes")
}
