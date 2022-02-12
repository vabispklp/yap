package handlers

import (
	"context"

	"github.com/vabispklp/yap/internal/app/service/model"
	storageModel "github.com/vabispklp/yap/internal/app/storage/model"
)

type ShortenerExpected interface {
	GetRedirectLink(ctx context.Context, id string) (*storageModel.ShortURL, error)
	GetUserURLs(ctx context.Context, userID string) ([]model.Shorten, error)
	AddRedirectLink(ctx context.Context, stringURL, userID string) (string, error)
	AddManyRedirectLink(ctx context.Context, shortenBatchItems []model.ShortenBatch, userID string) ([]model.ShortenBatchResult, error)

	Ping(ctx context.Context) error
}
