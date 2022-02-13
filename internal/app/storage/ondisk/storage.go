package ondisk

import (
	"context"
	"encoding/json"
	"github.com/vabispklp/yap/internal/app/storage"
	"github.com/vabispklp/yap/internal/app/storage/model"
	"io"
	"os"
	"sync"
)

type Storage struct {
	sync.Mutex
	filePath string

	writeFile *os.File
	readFile  *os.File

	encoder *json.Encoder
}

func NewStorage(filePath string) (storage.StorageExpected, error) {
	writeFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		return nil, err
	}

	readFile, err := os.OpenFile(filePath, os.O_RDONLY, 0777)
	if err != nil {
		return nil, err
	}

	return &Storage{
		filePath:  filePath,
		writeFile: writeFile,
		readFile:  readFile,
		encoder:   json.NewEncoder(writeFile),
	}, nil
}

func (s *Storage) Get(_ context.Context, id string) (*model.ShortURL, error) {
	return s.getByID(id)
}

func (s *Storage) Add(_ context.Context, shortURL model.ShortURL) error {
	savedURL, err := s.getByID(shortURL.ID)
	if err != nil {
		return err
	}
	if savedURL != nil {
		return nil
	}

	return s.encoder.Encode(shortURL)
}

func (s *Storage) GetByUser(_ context.Context, userID string) ([]model.ShortURL, error) {
	var (
		item model.ShortURL
	)

	s.Lock()
	defer s.Unlock()

	result := make([]model.ShortURL, 0)
	_, err := s.readFile.Seek(0, io.SeekStart)
	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(s.readFile)
	for decoder.More() {
		err = decoder.Decode(&item)
		if err != nil {
			return nil, err
		}

		if item.UserID == userID {
			result = append(result, item)
		}
	}

	return result, nil
}

func (s *Storage) getByID(id string) (*model.ShortURL, error) {
	var result, item *model.ShortURL

	s.Lock()
	defer s.Unlock()

	_, err := s.readFile.Seek(0, io.SeekStart)
	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(s.readFile)
	for decoder.More() {
		err = decoder.Decode(&item)
		if err != nil {
			return nil, err
		}

		if item.ID == id {
			result = item
			break
		}
	}

	return result, nil
}

func (s *Storage) AddMany(_ context.Context, _ []model.ShortURL) error {
	return nil
}

func (s *Storage) Close() error {
	err := s.readFile.Close()
	if err != nil {
		return err
	}

	return s.writeFile.Close()
}

func (s *Storage) Ping(_ context.Context) error {
	return nil
}

func (s *Storage) Delete(_ context.Context, _ []string, _ string) error {
	return nil
}
