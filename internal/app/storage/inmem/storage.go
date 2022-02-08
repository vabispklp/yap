package inmem

import (
	"context"
	"github.com/vabispklp/yap/internal/app/storage"
	"github.com/vabispklp/yap/internal/app/storage/model"
	"sync"
)

type Storage struct {
	sync.Mutex
	urlsMap map[string]model.ShortURL
}

func NewStorage() storage.StorageExpected {
	return &Storage{
		urlsMap: make(map[string]model.ShortURL),
	}
}

func (s *Storage) Get(_ context.Context, id string) (*model.ShortURL, error) {
	shortURL, ok := s.urlsMap[id]
	if !ok {
		return nil, nil
	}

	return &shortURL, nil
}

func (s *Storage) Add(_ context.Context, shortURL model.ShortURL) error {
	s.Lock()
	s.urlsMap[shortURL.ID] = shortURL
	s.Unlock()

	return nil
}

func (s *Storage) GetByUser(ctx context.Context, userID string) ([]model.ShortURL, error) {
	return nil, nil
}

func (s *Storage) Close() error {
	return nil
}

func (s *Storage) Ping(_ context.Context) error {
	return nil
}
