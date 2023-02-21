package cmd

import (
	"fmt"
	"github.com/howeyc/gopass"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"ockham-api/config"
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Setup your own configure file.",
	Long:  `Setup your own configure file.`,
	Run: func(cmd *cobra.Command, args []string) {
		if config.DbHost == "" {
			fmt.Print("Input DBS host (127.0.0.1): ")
			fmt.Scanln(&config.DbHost)
			fmt.Print("Input DBS port (3306): ")
			fmt.Scanln(&config.DbPort)
			if config.DbHost == "" {
				config.DbHost = "127.0.0.1"
				config.DbPort = 3306
			}
		}

		if config.DbSchema == "" {
			fmt.Print("Input DBS schema (ockham): ")
			fmt.Scanln(&config.DbSchema)
			if config.DbSchema == "" {
				config.DbSchema = "ockham"
			}
		}

		for {
			if config.DbUser == "" {
				fmt.Print("Input DBS username: ")
				fmt.Scanln(&config.DbUser)
			} else {
				break
			}
		}

		for {
			if config.DbPass == "" {
				fmt.Print("Input DBS password: ")
				pass, _ := gopass.GetPasswd()
				config.DbPass = string(pass)
			} else {
				break
			}
		}

		/*
			if config.DbCharset == "" {
				fmt.Print("Input client-to-DBS connection charset (utf8mb4): ")
				fmt.Scanln(&config.DbCharset)
				if config.DbCharset == "" {
					config.DbCharset = "utf8mb4"
				}
			}
		*/

		err := viper.WriteConfig()
		if err != nil {
			fmt.Println("write config err: ", err)
			return
		} else {
			fmt.Println("config wrote.")
		}
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)

	configureCmd.Flags().StringVarP(&config.DbHost, "db-host", "d", "", "Host of DBS to connect.")
	viper.BindPFlag("db.host", configureCmd.Flags().Lookup("db-host"))

	configureCmd.Flags().IntVarP(&config.DbPort, "db-port", "p", 3306, "Port of DBS to connect.")
	viper.BindPFlag("db.port", configureCmd.Flags().Lookup("db-port"))

	configureCmd.Flags().StringVarP(&config.DbSchema, "db-schema", "s", "", "Schema(database) of DBS to connect.")
	viper.BindPFlag("db.schema", configureCmd.Flags().Lookup("db-schema"))

	configureCmd.Flags().StringVarP(&config.DbUser, "db-user", "u", "", "Username of DBS to connect.")
	viper.BindPFlag("db.user", configureCmd.Flags().Lookup("db-user"))

	configureCmd.Flags().StringVarP(&config.DbPass, "db-pass", "k", "", "Password of DBS to connect.")
	viper.BindPFlag("db.pass", configureCmd.Flags().Lookup("db-pass"))

	configureCmd.Flags().StringVar(&config.DbCharset, "db-charset", "utf8mb4", "Encoding charset of client-to-DBS connections.")
	viper.BindPFlag("db.charset", configureCmd.Flags().Lookup("db-charset"))
}
