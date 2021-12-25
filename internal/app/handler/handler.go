package handler

import (
	"context"

	"github.com/vabispklp/yap/internal/app/model"
)

type Handler struct {
	storage Storage
}

func New(storage Storage) *Handler {
	return &Handler{
		storage: storage,
	}
}

type Storage interface {
	GetRedirectLink(ctx context.Context, path string) (*model.ShortURL, error)
	AddRedirectLink(ctx context.Context, shortURL *model.ShortURL) error
}
