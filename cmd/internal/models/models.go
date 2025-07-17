package models

import "time"

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

type HTTPError struct {
	Timestamp string `json:"timestamp"`
	Status    int    `json:"status"`
	Error     string `json:"error"`
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
