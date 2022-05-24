package cmd

import (
	"fmt"
	"github.com/spf13/viper"
	"ockham-api/config"

	"github.com/spf13/cobra"
)

// configureAuthCmd represents the configureAuthsession command
var configureAuthCmd = &cobra.Command{
	Use:   "auth",
	Short: "Configure auth session.",
	Long:  `Configure auth session.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("configure auth called")
	},
}

func init() {
	configureCmd.AddCommand(configureAuthCmd)

	configureAuthCmd.Flags().IntVar(&config.AuthSessionExpireSeconds, "session-expire-seconds", 7200, "Auth session expire seconds.")
	viper.BindPFlag("auth.session.expire_seconds", configureAuthCmd.Flags().Lookup("session-expire-seconds"))

	configureAuthCmd.Flags().IntVar(&config.AuthSessionMaximumRenewalTimes, "session-maximum-renewal-times", 3, "Auth session maximum renewal times.")
	viper.BindPFlag("auth.session.expire_seconds", configureAuthCmd.Flags().Lookup("session-maximum-renewal-times"))

	configureAuthCmd.Flags().StringVar(&config.AuthJwtIssuer, "jwt-issuer", "ockham-api", "Auth JWT issuer.")
	viper.BindPFlag("auth.jwt.issuer", configureAuthCmd.Flags().Lookup("jwt-issuer"))

	configureAuthCmd.Flags().StringVar(&config.AuthJwtSecret, "jwt-secret", "your-own-secret-here", "Auth JWT secret.")
	viper.BindPFlag("auth.jwt.secret", configureAuthCmd.Flags().Lookup("jwt-secret"))

	configureAuthCmd.Flags().IntVar(&config.AuthJwtExpireSeconds, "jwt-expire-seconds", 1800, "Auth JWT expire seconds.")
	viper.BindPFlag("auth.jwt.expire_seconds", configureAuthCmd.Flags().Lookup("jwt-expire-seconds"))

	configureAuthCmd.Flags().IntVar(&config.AuthSignatureTimestampToleranceSeconds, "signature-timestamp-tolerance-seconds", 60, "Auth signature tolerance seconds.")
	viper.BindPFlag("auth.signature.timestamp-tolerance-seconds", configureAuthCmd.Flags().Lookup("signature-timestamp-tolerance-seconds"))

	configureAuthCmd.Flags().IntVar(&config.AuthSignatureBodyDigestTruncateBytes, "signature-body-digest-truncate-bytes", 128, "Auth signature body digest truncate bytes.")
	viper.BindPFlag("auth.signature.body-digest-truncate-bytes", configureAuthCmd.Flags().Lookup("signature-body-digest-truncate-bytes"))
}
