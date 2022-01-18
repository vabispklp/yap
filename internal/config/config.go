package config

import (
	"net/url"

	"github.com/caarlos0/env/v6"
)

type config struct {
	ServerAddr      string  `env:"SERVER_ADDRESS" envDefault:"localhost:8080"`
	BaseURL         url.URL `env:"BASE_URL" envDefault:"http://localhost:8080"`
	FileStoragePath string  `env:"FILE_STORAGE_PATH" envDefault:"tmp"`
}

func InitConfig() (*config, error) {
	var cfg config
	if err := env.Parse(&cfg); err != nil {
		return nil, err
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
