package shortener

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"net/url"

	"github.com/vabispklp/yap/internal/app/model"
	"github.com/vabispklp/yap/internal/app/storage"
)

type Shortener struct {
	storage storage.StorageExpected

	baseURL url.URL
}

func NewShortener(storage storage.StorageExpected, baseURL url.URL) (*Shortener, error) {
	if storage == nil {
		return nil, ErrNilPointerStorage
	}

	return &Shortener{
		storage: storage,
		baseURL: baseURL,
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

	u := s.baseURL
	u.Path = resultPath

	err := s.storage.AddRedirectLink(ctx, model.ShortURL{
		ID:          resultPath,
		OriginalURL: stringURL,
	})
	if err != nil {
		return "", err
	}

	return u.String(), nil
}
