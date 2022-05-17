package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// manageCmd represents the admin command
var manageCmd = &cobra.Command{
	Use:   "manage",
	Short: "Admin tools",
	Long:  `Admin tools that can manage your super user or other functions.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("manage called")
	},
}

func init() {
	rootCmd.AddCommand(manageCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// adminCmd.PersistentFlags().String("foo", "", "A help for foo")
	//adminCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//adminCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	manageCmd.Flags().StringVarP(&dbHost, "db-host", "d", "localhost", "Database host.")
	manageCmd.Flags().StringVarP(&dbSchema, "db-schema", "s", "ockham", "Database schema.")
	manageCmd.Flags().StringVarP(&dbUser, "db-user", "u", "root", "Database username.")
	manageCmd.Flags().StringVarP(&dbPass, "db-pass", "p", "123456", "Database password.")
}
