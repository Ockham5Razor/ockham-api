package cmd

import (
	"fmt"
	"github.com/spf13/viper"
	"ockham-api/config"

	"github.com/spf13/cobra"
)

// configureEmailCmd represents the configureEmail command
var configureEmailCmd = &cobra.Command{
	Use:   "email",
	Short: "Configure email post account.",
	Long:  `Configure email post account.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("configureEmail called")
	},
}

func init() {
	configureCmd.AddCommand(configureEmailCmd)

	configureEmailCmd.Flags().StringVar(&config.EmailUser, "user", "tommy1995@163.com", "Email username.")
	viper.BindPFlag("email.user", configureEmailCmd.Flags().Lookup("user"))

	configureEmailCmd.Flags().StringVar(&config.EmailPass, "pass", "t0mmy123", "Email password.")
	viper.BindPFlag("email.pass", configureEmailCmd.Flags().Lookup("pass"))

	configureEmailCmd.Flags().StringVar(&config.EmailHost, "host", "smtp.163.com", "Email host.")
	viper.BindPFlag("email.host", configureEmailCmd.Flags().Lookup("host"))

	configureEmailCmd.Flags().IntVar(&config.EmailPort, "port", 465, "Email host port.")
	viper.BindPFlag("email.port", configureEmailCmd.Flags().Lookup("port"))

	configureEmailCmd.Flags().StringVar(&config.EmailSign, "sign", "ockham-api", "Email sign.")
	viper.BindPFlag("email.sign", configureEmailCmd.Flags().Lookup("sign"))

}
