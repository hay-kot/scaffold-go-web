// Package config contains the shared configuration for startup code.
package config

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/ardanlabs/conf/v3"
)

type Mode string

func (m Mode) String() string {
	return string(m)
}

func (m Mode) IsDevelopment() bool {
	return m == "development"
}

func (m Mode) IsProduction() bool {
	return !m.IsDevelopment()
}

type Config struct {
	Mode     Mode   `conf:"default:development"`
	LogLevel string `conf:"default:info"`
	Web      WebConfig
	Postgres PostgresConfig
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

type PostgresSSLMode string

func (p PostgresSSLMode) String() string {
	lower := strings.ToLower(string(p))

	switch lower {
	case "disable":
		return "disable"
	case "verify-full":
		return "verify-full"
	default:
		return "require"
	}
}

type PostgresConfig struct {
	User         string          `conf:"default:postgres"`
	Password     string          `conf:"default:postgres,mask"`
	Host         string          `conf:"default:localhost"`
	Port         string          `conf:"default:5432"`
	Name         string          `conf:"default:goapi"`
	MaxIdleConns int             `conf:"default:2"`
	MaxOpenConns int             `conf:"default:0"`
	SSLMode      PostgresSSLMode `conf:"default:verify-full"`
}

// NewFromCLI parses the CLI/Config file and returns a Config struct. If the file argument is an empty string, the
// file is not read. If the file is not empty, the file is read and the Config struct is returned.
func NewFromCLI() (*Config, error) {
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
