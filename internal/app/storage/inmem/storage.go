package inmem

import (
	"context"
	"github.com/vabispklp/yap/internal/app/storage"
	"github.com/vabispklp/yap/internal/app/storage/model"
	"sync"
)

// Storage хранилище реализованное в памяти
type Storage struct {
	sync.Mutex
	urlsMap map[string]model.ShortURL
}

// NewStorage создает Storage
func NewStorage() storage.StorageExpected {
	return &Storage{
		urlsMap: make(map[string]model.ShortURL),
	}
}

// Get отдает сокращенную ссылку
func (s *Storage) Get(_ context.Context, id string) (*model.ShortURL, error) {
	shortURL, ok := s.urlsMap[id]
	if !ok {
		return nil, nil
	}

	return &shortURL, nil
}

// Add добавленяет сокращенную ссылки
func (s *Storage) Add(_ context.Context, shortURL model.ShortURL) error {
	s.Lock()
	s.urlsMap[shortURL.ID] = shortURL
	s.Unlock()

	return nil
}

// GetByUser отдает все ссылки пользователей
func (s *Storage) GetByUser(_ context.Context, _ string) ([]model.ShortURL, error) {
	return nil, nil
}

// AddMany добавляет несколько ссылок одновременно
func (s *Storage) AddMany(_ context.Context, _ []model.ShortURL) error {
	return nil
}

// Close закрывает соединение с хранилищем
func (s *Storage) Close() error {
	return nil
}

// Ping пингует хранилище
func (s *Storage) Ping(_ context.Context) error {
	return nil
}

// Delete удаляет сокращенные ссылки
func (s *Storage) Delete(_ context.Context, _ []string, _ string) error {
	return nil
}
