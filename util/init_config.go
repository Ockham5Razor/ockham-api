package util

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

func InitConfig(cfgFile string) {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home + "/.ockham-api/")
		viper.SetConfigName("configs")
		viper.SetConfigType("yaml")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			fmt.Println("config file not found.")
		} else {
			// Config file was found but another error was produced
			fmt.Println("config file read other error: ", err.Error())
		}
		os.Exit(2)
	} else {
		fmt.Println("using config file: ", viper.ConfigFileUsed())
	}
	// Config file found and successfully parsed
}
