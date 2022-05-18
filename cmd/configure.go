package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	dbHost    string
	dbPort    int
	dbSchema  string
	dbUser    string
	dbPass    string
	dbCharset string
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Setup your own configure file.",
	Long:  `Setup your own configure file.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := viper.SafeWriteConfig()
		if err != nil {
			fmt.Println("write config err: ", err)
			return
		} else {
			fmt.Println("wrote")
		}
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configureCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configureCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	//configureCmd.Flags().StringVarP(&dbHost, "profile", "", "default", "setting profile")
	configureCmd.Flags().StringVarP(&dbHost, "db-host", "d", "", "Host of DBS to connect.")
	configureCmd.Flags().IntVarP(&dbPort, "db-port", "p", 3306, "Port of DBS to connect.")
	configureCmd.Flags().StringVarP(&dbSchema, "db-schema", "s", "", "Schema(database) of DBS to connect.")
	configureCmd.Flags().StringVarP(&dbUser, "db-user", "u", "", "Username of DBS to connect.")
	configureCmd.Flags().StringVarP(&dbPass, "db-pass", "k", "", "Password of DBS to connect.")
	configureCmd.Flags().StringVar(&dbCharset, "db-charset", "utf8mb4", "Encoding charset of client-to-DBS connections.")
	viper.BindPFlag("db.host", configureCmd.Flags().Lookup("db-host"))
	viper.BindPFlag("db.port", configureCmd.Flags().Lookup("db-port"))
	viper.BindPFlag("db.schema", configureCmd.Flags().Lookup("db-schema"))
	viper.BindPFlag("db.user", configureCmd.Flags().Lookup("db-user"))
	viper.BindPFlag("db.pass", configureCmd.Flags().Lookup("db-pass"))
	viper.BindPFlag("db.charset", configureCmd.Flags().Lookup("db-charset"))
}
