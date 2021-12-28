package shortener

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/vabispklp/yap/internal/app/storage"

	"github.com/vabispklp/yap/internal/app/model"
)

const (
	resultURLPattern = "http://localhost:8080/%s"
)

type Shortener struct {
	storage storage.ShortenerExpected
}

func NewShortener(storage storage.ShortenerExpected) (*Shortener, error) {
	if storage == nil {
		return nil, ErrNilPointerStorage
	}

	return &Shortener{
		storage: storage,
	}, nil
}

func (s *Shortener) GetRedirectLink(ctx context.Context, id string) (*model.ShortURL, error) {
	result, err := s.storage.GetRedirectLink(ctx, id)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, ErrNotFound
	}

	return result, nil
}

func (s *Shortener) AddRedirectLink(ctx context.Context, stringURL string) (string, error) {
	hash := md5.Sum([]byte(stringURL))
	resultPath := hex.EncodeToString(hash[:])

	resultURL := fmt.Sprintf(resultURLPattern, resultPath)

	err := s.storage.AddRedirectLink(ctx, &model.ShortURL{
		ID:          resultPath,
		OriginalURL: stringURL,
	})
	if err != nil {
		return "", err
	}

	return resultURL, nil
}
