package handlers

import (
	"context"

	"github.com/vabispklp/yap/internal/app/model"
)

type Handler struct {
	service ShortenerExpected
}

func NewHandler(service ShortenerExpected) (*Handler, error) {
	if service == nil {
		return nil, ErrNilPointerService
	}

	return &Handler{
		service: service,
	}, nil
}

type ShortenerExpected interface {
	GetRedirectLink(ctx context.Context, id string) (*model.ShortURL, error)
	AddRedirectLink(ctx context.Context, stringURL string) (string, error)
}
