package config

import "net/url"

type ConfigExpected interface {
	GetServerAddr() string
	GetBaseURL() url.URL
	GetFileStoragePath() string
}
