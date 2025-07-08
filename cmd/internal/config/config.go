package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/siestaw/laterna/server/cmd/internal/models"
)

func LoadConfig(path string) *models.Config {
	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("An error occured while reading the config file:", err)
		os.Exit(1)
	}

	var cfg models.Config
	err = json.Unmarshal(file, &cfg)
	if err != nil {
		fmt.Println("An error occured while parsing the config file:", err)
		os.Exit(1)
	}

	return &cfg
}

var AppConfig *models.Config
