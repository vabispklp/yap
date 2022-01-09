package storage

import (
	"context"

	"github.com/vabispklp/yap/internal/app/model"
)

type StorageExpected interface {
	GetRedirectLink(ctx context.Context, id string) (*model.ShortURL, error)
	AddRedirectLink(ctx context.Context, shortURL model.ShortURL) error
}
