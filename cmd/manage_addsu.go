package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// adminaddsuCmd represents the adminaddsu command
var adminaddsuCmd = &cobra.Command{
	Use:   "addsu",
	Short: "Add super user",
	Long:  `Add super user`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("addsu called")
	},
}

func init() {
	manageCmd.AddCommand(adminaddsuCmd)
}
