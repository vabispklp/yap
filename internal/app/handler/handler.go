package handler

import (
	"context"

	"github.com/vabispklp/yap/internal/app/model"
)

type Handler struct {
	storage Repository
}

func New(storage Repository) *Handler {
	return &Handler{
		storage: storage,
	}
}

type Repository interface {
	GetRedirectLink(ctx context.Context, path string) (*model.ShortURL, error)
	AddRedirectLink(ctx context.Context, shortURL *model.ShortURL) error
}
