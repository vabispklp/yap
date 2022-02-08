package shortener

import (
	"net/url"

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
