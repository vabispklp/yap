package storage

import (
	"context"
	"errors"
	"sync"

	"github.com/vabispklp/yap/internal/app/model"
)

var errNotFound = errors.New("original URL not found")

type Storage struct {
	urlsMap map[string]*model.ShortURL
}

func New() *Storage {
	return &Storage{
		urlsMap: make(map[string]*model.ShortURL),
	}
}

func (r *Storage) GetRedirectLink(ctx context.Context, path string) (*model.ShortURL, error) {
	shortURL, ok := r.urlsMap[path]
	if !ok {
		return nil, errNotFound
	}

	return shortURL, nil
}

func (r *Storage) AddRedirectLink(ctx context.Context, shortURL *model.ShortURL) error {
	mu := sync.Mutex{}
	mu.Lock()
	r.urlsMap[shortURL.Path] = shortURL
	mu.Unlock()

	return nil
}
