package util

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
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
		cfgFile = filepath.Join(home, ".ockham-api", "configs.yaml")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Config file not found or read error, should we create or override config file: `%v`? (Y/n): ", cfgFile)
		var shouldOverride = "Y"
		_, scanInputErr := fmt.Scanln(&shouldOverride)
		if scanInputErr != nil {
			log.Println("Invalid input, using default value [Y].")
		}
		if shouldOverride == "Y" || shouldOverride == "y" || shouldOverride == "" {
			fmt.Printf("Creating or overrideing config file: `%v`...\n", cfgFile)
			dir := filepath.Dir(cfgFile)
			mkdirErr := os.MkdirAll(dir, 0755)
			if mkdirErr != nil {
				fmt.Printf("Error creating or overrideing config file: `%v`! Error is: %v\n", cfgFile, mkdirErr)
				os.Exit(3)
			}
			f, createFileErr := os.Create(cfgFile)
			defer f.Close()
			if createFileErr != nil {
				fmt.Printf("Error creating or overrideing config file: `%v`! Error is: %v\n", cfgFile, createFileErr)
				os.Exit(4)
			}
		} else {
			fmt.Printf("Canceled creating or overrideing config file: `%v`.\n", cfgFile)
			os.Exit(2)
		}
	} else {
		fmt.Println("Using config file: ", cfgFile)
	}
	// Config file found and successfully parsed
}
