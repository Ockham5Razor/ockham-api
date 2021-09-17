package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
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
}

var conf Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		data := readConfigFile()
		err := yaml.Unmarshal(data, &conf)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
	})
	return &conf
}

func readConfigFile() []byte {
	bytes, err := ioutil.ReadFile("configs.yaml")
	if err != nil {
		log.Fatal(err)
	}
	return bytes
}
