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
	Status    int    `json:"status"`
	Success   bool   `json:"success"`
	Error     string `json:"error,omitempty"`
	Timestamp string `json:"timestamp"`
	Data      any    `json:"data,omitempty"`
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

type DeleteData struct {
	Deleted int `json:"deleted"`
}

type CreateData struct {
	Created int `json:"created"`
}
