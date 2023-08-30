package util

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"ockham-api/config"
	"ockham-api/database"
	"os"
)

const ConfigFilePath = "C:\\Users\\hera\\ockham-api\\configs.yaml"

func InitDatabase() *gorm.DB {
	path, err := os.Executable()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)

	InitConfig(ConfigFilePath) // specify your config file
	config.FillParams()
	return database.Init()
}
