package web

import (
	"errors"
	"fmt"
	"os"

	"github.com/ardanlabs/conf/v3"
)

const (
	ModeDevelopment = "development"
	ModeProduction  = "production"
)

type Config struct {
	Mode     string `conf:"default:development"`
	LogLevel string `conf:"default:info"`
	Web      WebConfig
}

type WebConfig struct {
	Host           string `conf:"default:0.0.0.0"`
	Port           string `conf:"default:8080"`
	AllowedOrigins string `conf:"default:https://*,http://*"`
	TLSCert        string
	TLSKey         string
	IdleTimeout    int `conf:"default:30"`
	ReadTimeout    int `conf:"default:10"`
	WriteTimeout   int `conf:"default:10"`
}

// ConfigFromCLI parses the CLI/Config file and returns a Config struct. If the file argument is an empty string, the
// file is not read. If the file is not empty, the file is read and the Config struct is returned.
func ConfigFromCLI() (*Config, error) {
	var cfg Config
	const prefix = "API"

	help, err := conf.Parse(prefix, &cfg)
	if err != nil {
		if errors.Is(err, conf.ErrHelpWanted) {
			fmt.Println(help)
			os.Exit(0)
		}
		return &cfg, fmt.Errorf("parsing config: %w", err)
	}

	return &cfg, nil
}
