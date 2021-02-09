package config

import (
	"gopkg.in/yaml.v3"
)

type config_yaml struct {
	PublicKey string `yaml:public_key`
	PrivateKey string `yaml:private_key`
}

type config struct {
	config config_yaml
}
var instance *config

func GetInstance(path ...string) *config {
	if instance == nil {
		instance = &config{config:readConfig(path[0])}
	}
	return instance
}

func readConfig(string path) config_yaml {

}
