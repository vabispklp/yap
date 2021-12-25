package storage

import (
	"context"
	"errors"
	"sync"

	"github.com/vabispklp/yap/internal/app/model"
)

var errNotFound = errors.New("original URL not found")

type Repository struct {
	urlsMap map[string]*model.ShortURL
}

func New() *Repository {
	return &Repository{
		urlsMap: make(map[string]*model.ShortURL),
	}
}

func (r *Repository) GetRedirectLink(ctx context.Context, path string) (*model.ShortURL, error) {
	shortURL, ok := r.urlsMap[path]
	if !ok {
		return nil, errNotFound
	}

	return shortURL, nil
}

func (r *Repository) AddRedirectLink(ctx context.Context, shortURL *model.ShortURL) error {
	mu := sync.Mutex{}
	mu.Lock()
	r.urlsMap[shortURL.Path] = shortURL
	mu.Unlock()

	return nil
}
