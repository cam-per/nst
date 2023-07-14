package config

import (
	"encoding/json"
	"os"
)

type xvfbCfg struct {
	Enabled bool `json:"enabled,omitempty"`
}

type seleniumCfg struct {
	Port       int        `json:"port,omitempty"`
	ServerPath string     `json:"server_path,omitempty"`
	Drivers    driversCfg `json:"drivers,omitempty"`
	XVFB       xvfbCfg    `json:"xvfb,omitempty"`
}

type driversCfg struct {
	ChromePath string `json:"chrome_path,omitempty"`
	GeckoPath  string `json:"gecko_path,omitempty"`
}

type conf struct {
	MimeType string      `json:"mime_type,omitempty"`
	Selenium seleniumCfg `json:"selenium,omitempty"`
}

var cfg conf

var (
	MimeType = &cfg.MimeType
	Selenium = &cfg.Selenium
)

func Load(path string) error {
	if path == "" {
		path = "config.json"
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &cfg)
}
