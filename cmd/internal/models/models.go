package models

import (
	"time"
)

type HTTPConfig struct {
	AdminToken string  `json:"adminToken"`
	Port       int     `json:"port"`
	Cooldown   float64 `json:"cooldown"`
}


type Config struct {
	HTTP           HTTPConfig `json:"http"`
	FileLogging    bool       `json:"fileLogging"`
	VerboseLogging bool       `json:"verboseLogging"`
}

type HTTPResponse struct {
	Timestamp string `json:"timestamp"`
	Status    int    `json:"status"`
	Text      string `json:"text"`
	Message   string `json:"message"`
	Path      string `json:"path"`
}

type LampState struct {
	ID        int       `json:"id"`
	Color     string    `json:"color"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LampUpdateRequest struct {
	Color string `json:"color"`
}

type ControllerRequests struct {
	ID int `json:"ID"`
}