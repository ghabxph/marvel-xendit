package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type config_yaml struct {
	PublicKey string `yaml:"public_key"`
	PrivateKey string `yaml:"private_key"`
}

type Config struct {
	config *config_yaml
}
var instance *Config

func GetInstance(path ...string) *Config {
	if instance == nil {
		instance = &Config{config:readConfig(path[0])}
	}
	return instance
}

func readConfig(path string) *config_yaml {

	// Gets file. Terminates the program if config path does
	// not exist.
	file, err := ioutil.ReadFile(path)
	panicIfError(err)

	conf_yaml := &config_yaml{}

	yaml.Unmarshal(file, conf_yaml)

	return conf_yaml
}

func panicIfError(e error) {
	// Panics if there's error.
	if e != nil {
		panic(e)
	}
}

// Returns the marvel public key
func (c *Config) PublicKey() string {
	return c.config.PublicKey
}

// Returns the marvel private key
func (c *Config) PrivateKey() string {
	return c.config.PrivateKey
}
