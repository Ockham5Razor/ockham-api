package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// AminCmd represents the admin command
var AminCmd = &cobra.Command{
	Use:   "admin",
	Short: "Admin tools",
	Long:  `Admin tools that can manage your super user or other functions.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("admin called")
	},
}

func init() {
	rootCmd.AddCommand(AminCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// adminCmd.PersistentFlags().String("foo", "", "A help for foo")
	//adminCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//adminCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	serveCmd.Flags().StringVarP(&dbHost, "db-host", "d", "localhost", "Database host.")
	serveCmd.Flags().StringVarP(&dbSchema, "db-schema", "s", "ockham", "Database schema.")
	serveCmd.Flags().StringVarP(&dbUser, "db-user", "u", "root", "Database username.")
	serveCmd.Flags().StringVarP(&dbPass, "db-pass", "p", "123456", "Database password.")
}
