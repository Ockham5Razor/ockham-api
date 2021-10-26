package config

import (
	"gopkg.in/yaml.v2"
	"log"
	"sync"
)

type Config struct {
	Db struct {
		Host    string
		Port    int
		User    string
		Pass    string
		Schema  string
		Charset string
	}
	Email struct {
		User string
		Pass string
		Host string
		Port int
		Sign string
	}
	EmailValidation struct {
		ExpireDuration string
	}
	Auth struct {
		JwtSecret string
	}
}

var conf Config
var confOnce sync.Once

func GetConfig() Config {
	confOnce.Do(func() {
		data := readConfigFile("configs.yaml")
		err := yaml.Unmarshal(data, &conf)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
	})
	return conf
}
