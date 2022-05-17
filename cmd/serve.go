package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"ockham-api/config"
)

var configFile string
var dbHost string
var dbSchema string
var dbUser string
var dbPass string
var listen string

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run service",
	Long:  `Run service.`,
	Run: func(cmd *cobra.Command, args []string) {
		config.ConfFilePath = configFile
		fmt.Println(configFile)
		//run.Main()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().StringVarP(&configFile, "config-file", "f", "configs.yaml", "Specify a config file, which will override all configs.")
	serveCmd.Flags().StringVarP(&dbHost, "db-host", "d", "localhost", "Database host.")
	serveCmd.Flags().StringVarP(&dbSchema, "db-schema", "s", "ockham", "Database schema.")
	serveCmd.Flags().StringVarP(&dbUser, "db-user", "u", "root", "Database username.")
	serveCmd.Flags().StringVarP(&dbPass, "db-pass", "p", "123456", "Database password.")
	serveCmd.Flags().StringVarP(&listen, "listen", "l", "0.0.0.0:8080", "Listen Address.")
}
