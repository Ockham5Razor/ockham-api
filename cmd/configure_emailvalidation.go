package cmd

import (
	"fmt"
	"github.com/spf13/viper"
	"ockham-api/config"

	"github.com/spf13/cobra"
)

// configureEmailValidationCmd represents the configureEmailValidation command
var configureEmailValidationCmd = &cobra.Command{
	Use:   "emailvalidation",
	Short: "Configure email validation.",
	Long:  `Configure email validation.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("configure emailvalidation called")
	},
}

func init() {
	configureCmd.AddCommand(configureEmailValidationCmd)

	configureCmd.Flags().StringVar(&config.EmailValidationExpireDuration, "expire-duration", "30m", "Email validation expire duration.")
	viper.BindPFlag("email_validation.expire_duration", configureCmd.Flags().Lookup("expire-duration"))
}
