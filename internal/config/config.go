package config

import (
	"flag"
	"net/url"

	"github.com/caarlos0/env/v6"
)

// Сonfig с environments
type Сonfig struct {
	ServerAddr      string  `env:"SERVER_ADDRESS"`
	BaseURL         url.URL `env:"BASE_URL"`
	FileStoragePath string  `env:"FILE_STORAGE_PATH"`
	DatabaseDSN     string  `env:"DATABASE_DSN"`
	PprofAddr       string  `env:"PPROF_ADDRESS" envDefault:":8084"`
}

var flagConfig struct {
	ServerAddr      string
	BaseURL         string
	FileStoragePath string
	DatabaseDSN     string
}

func init() {
	flag.StringVar(&flagConfig.ServerAddr, "a", "localhost:8080", "host:port to listen")
	flag.StringVar(&flagConfig.BaseURL, "b", "http://localhost:8080", "base url shorten URL")
	flag.StringVar(&flagConfig.FileStoragePath, "f", "tmp", "path to storage file")
	flag.StringVar(&flagConfig.DatabaseDSN, "d", "host=localhost port=5432 user=user dbname=shortener sslmode=disable", "database dsn")
	flag.Parse()
}

// GetConfig возвращает проинициализованный Сonfig
func GetConfig() (*Сonfig, error) {
	var cfg Сonfig

	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	if cfg.ServerAddr == "" {
		cfg.ServerAddr = flagConfig.ServerAddr
	}

	if cfg.BaseURL.String() == "" {
		baseURL, err := url.Parse(flagConfig.BaseURL)
		if err != nil {
			return nil, err
		}
		cfg.BaseURL = *baseURL
	}

	if cfg.FileStoragePath == "" {
		cfg.FileStoragePath = flagConfig.FileStoragePath
	}

	if cfg.DatabaseDSN == "" {
		cfg.DatabaseDSN = flagConfig.DatabaseDSN
	}

	return &cfg, nil
}
