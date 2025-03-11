package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Features map[string]bool `yaml:"features"`
}

func LoadConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func IsFeatureEnable(config *Config, feature string) bool {
	enabled, exists := config.Features[feature]

	return exists && enabled
}

func main() {
	config, err := LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Erro ao carregar configuração: %v", err)
	}

	featureName := "new_dashboard"
	if IsFeatureEnable(config, featureName) {
		fmt.Printf("A feature '%s' está ativada!\n", featureName)
	} else {
		fmt.Printf("A feature '%s' está desativada.\n", featureName)
	}

}
