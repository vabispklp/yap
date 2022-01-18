package config

import (
	"flag"
	"net/url"

	"github.com/caarlos0/env/v6"
)

type config struct {
	ServerAddr      string  `env:"SERVER_ADDRESS"`
	BaseURL         url.URL `env:"BASE_URL"`
	FileStoragePath string  `env:"FILE_STORAGE_PATH"`
}

var flagConfig struct {
	ServerAddr      string
	BaseURL         string
	FileStoragePath string
}

func init() {
	flag.StringVar(&flagConfig.ServerAddr, "a", "localhost:8080", "host:port to listen")
	flag.StringVar(&flagConfig.BaseURL, "b", "http://localhost:8080", "base url shorten URL")
	flag.StringVar(&flagConfig.FileStoragePath, "f", "tmp", "path to storage file")
	flag.Parse()
}

func InitConfig() (*config, error) {
	var cfg config

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

	return &cfg, nil
}

func (c *config) GetServerAddr() string {
	return c.ServerAddr
}

func (c *config) GetBaseURL() url.URL {
	return c.BaseURL
}

func (c *config) GetFileStoragePath() string {
	return c.FileStoragePath
}
