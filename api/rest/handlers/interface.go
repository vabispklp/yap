package handlers

import (
	"context"

	"github.com/vabispklp/yap/internal/app/model"
)

type ShortenerExpected interface {
	GetRedirectLink(ctx context.Context, id string) (*model.ShortURL, error)
	AddRedirectLink(ctx context.Context, stringURL string) (string, error)

	Ping(ctx context.Context) error
}
