package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Port     int `json:"port"`
	Database struct {
		User     string `json:"user"`
		Password string `json:"password"`
		DBName   string `json:"dbname"`
	} `json:"database"`
}

func loadConfig(filename string) (Config, error) {
	var config Config
	file, err := os.Open(filename)
	if err != nil {
		return config, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return config, err
	}
	return config, nil
}

func main() {
	config, err := loadConfig("config.json")
	if err != nil {
		fmt.Println("Erro ao carregar cofiguracoes: ", err)
		return
	}

	fmt.Println(config)

}
