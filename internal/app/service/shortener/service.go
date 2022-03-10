package shortener

import (
	"net/url"

	"github.com/vabispklp/yap/internal/app/storage"
)

// Shortener стуктура бизнес логики
type Shortener struct {
	storage storage.StorageExpected

	baseURL url.URL
}

// NewShortener создает стуркутуру Shortener
func NewShortener(storage storage.StorageExpected, baseURL url.URL) (*Shortener, error) {
	if storage == nil {
		return nil, ErrNilPointerStorage
	}

	return &Shortener{
		storage: storage,
		baseURL: baseURL,
	}, nil
}
