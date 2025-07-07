package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type HTTPConfig struct {
	AdminToken string `json:"adminToken"`
	Port       int    `json:"port"`
}

type Config struct {
	HTTP           HTTPConfig `json:"http"`
	FileLogging    bool       `json:"fileLogging"`
	VerboseLogging bool       `json:"verboseLogging"`
}

func LoadConfig(path string) *Config {
	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("An error occured while reading the config file:", err)
		os.Exit(1)
	}

	var cfg Config
	err = json.Unmarshal(file, &cfg)
	if err != nil {
		fmt.Println("An error occured while parsing the config file:", err)
		os.Exit(1)
	}

	return &cfg
}

var AppConfig *Config
