package apiserver

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	BindAddr string `yaml:"bind_addr"`
}

func NewConfig(cfgPath string) *Config {
	yamlFile, err := os.ReadFile(cfgPath)
	if err != nil {
		log.Println("Error reading YAML file. Loading default config.")
		return &Config{
			BindAddr: ":8080",
		}
	}

	var cfg Config
	err = yaml.Unmarshal(yamlFile, &cfg)
	if err != nil {
		log.Println("Error unmarshalling YAML. Loading default config.")
		return &Config{
			BindAddr: ":8080",
		}
	}

	return &cfg
}
