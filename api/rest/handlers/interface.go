package handlers

import (
	"context"

	"github.com/vabispklp/yap/internal/app/service/model"
	storageModel "github.com/vabispklp/yap/internal/app/storage/model"
)

type ShortenerExpected interface {
	GetRedirectLink(ctx context.Context, id string) (*storageModel.ShortURL, error)
	GetUserURLs(ctx context.Context, userID string) ([]model.ShortenItemResponse, error)
	AddRedirectLink(ctx context.Context, stringURL, userID string) (string, error)
	AddManyRedirectLink(ctx context.Context, shortenBatchItems []model.ShortenBatchRequest, userID string) ([]model.ShortenBatchResponse, error)

	Ping(ctx context.Context) error
}
