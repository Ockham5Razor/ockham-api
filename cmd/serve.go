package cmd

import (
	"github.com/spf13/cobra"
	"ockham-api/run"
)

var listen string

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run service",
	Long:  `Run ockham API service.`,
	Run: func(cmd *cobra.Command, args []string) {
		//config.ConfFilePath = configFile
		run.Main(listen)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().StringVarP(&listen, "listen", "l", "0.0.0.0:8080", "Listen Address.")
}
