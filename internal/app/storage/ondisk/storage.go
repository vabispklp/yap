package ondisk

import (
	"context"
	"encoding/json"
	"io"
	"os"
	"sync"

	"github.com/vabispklp/yap/internal/app/model"
)

type Storage struct {
	filePath string

	mu sync.Mutex

	writeFile *os.File
	readFile  *os.File

	encoder *json.Encoder
}

func New(filePath string) (*Storage, error) {
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
		mu:        sync.Mutex{},
		writeFile: writeFile,
		readFile:  readFile,
		encoder:   json.NewEncoder(writeFile),
	}, nil
}

func (s *Storage) GetRedirectLink(_ context.Context, id string) (*model.ShortURL, error) {
	return s.getByID(id)
}

func (s *Storage) AddRedirectLink(_ context.Context, shortURL model.ShortURL) error {
	savedURL, err := s.getByID(shortURL.ID)
	if err != nil {
		return err
	}
	if savedURL != nil {
		return nil
	}

	return s.encoder.Encode(shortURL)
}

func (s *Storage) getByID(id string) (*model.ShortURL, error) {
	var result, item *model.ShortURL

	s.mu.Lock()
	defer s.mu.Unlock()

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

// todo не знаю как лучше вынести в main
func (s *Storage) Close() error {
	err := s.readFile.Close()
	if err != nil {
		return err
	}

	return s.writeFile.Close()
}
