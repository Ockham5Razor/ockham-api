package config

import (
	"io/ioutil"
	"log"
)

func readConfigFile(configFilePath string) []byte {
	bytes, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		log.Fatal(err)
	}
	return bytes
}

