package config

import (
	"gopkg.in/yaml.v2"
	"log"
	"sync"
)

var testEnvConfOnce sync.Once

var testEnvConf map[interface{}]map[interface{}]string

func GetTestEnvConfig() map[interface{}]map[interface{}]string {
	testEnvConfOnce.Do(func() {
		data := readConfigFile("test_env.configs.yaml")
		err := yaml.Unmarshal(data, &testEnvConf)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
	})
	return testEnvConf
}
