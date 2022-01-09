package inmem

import (
	"context"
	"sync"

	"github.com/vabispklp/yap/internal/app/model"
)

type Storage struct {
	mu      sync.Mutex
	urlsMap map[string]model.ShortURL
}

func New() *Storage {
	return &Storage{
		urlsMap: make(map[string]model.ShortURL),
	}
}

func (s *Storage) GetRedirectLink(ctx context.Context, id string) (*model.ShortURL, error) {
	shortURL, ok := s.urlsMap[id]
	if !ok {
		return nil, nil
	}

	return &shortURL, nil
}

func (s *Storage) AddRedirectLink(ctx context.Context, shortURL model.ShortURL) error {
	s.mu.Lock()
	s.urlsMap[shortURL.ID] = shortURL
	s.mu.Unlock()

	return nil
}
