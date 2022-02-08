package handlers

import (
	"context"

	"github.com/vabispklp/yap/internal/app/service/model"
	storageModel "github.com/vabispklp/yap/internal/app/storage/model"
)

type ShortenerExpected interface {
	GetRedirectLink(ctx context.Context, id string) (*storageModel.ShortURL, error)
	AddRedirectLink(ctx context.Context, stringURL, userID string) (string, error)
	GetUserURLs(ctx context.Context, userID string) ([]model.Shorten, error)

	Ping(ctx context.Context) error
}
