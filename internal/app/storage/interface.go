package storage

import (
	"context"
	"github.com/vabispklp/yap/internal/app/storage/model"
)

// StorageExpected интерфейс хранилища
type StorageExpected interface {
	Get(ctx context.Context, id string) (*model.ShortURL, error)
	GetByUser(ctx context.Context, userID string) ([]model.ShortURL, error)
	Add(ctx context.Context, shortURL model.ShortURL) error
	AddMany(ctx context.Context, shortURLs []model.ShortURL) error

	Delete(ctx context.Context, ids []string, userID string) error

	Ping(ctx context.Context) error
	Close() error
}
