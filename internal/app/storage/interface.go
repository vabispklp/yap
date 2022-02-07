package storage

import (
	"context"

	"github.com/vabispklp/yap/internal/app/model"
)

type StorageExpected interface {
	Get(ctx context.Context, id string) (*model.ShortURL, error)
	Add(ctx context.Context, shortURL model.ShortURL) error

	Ping(ctx context.Context) error
	Close() error
}
