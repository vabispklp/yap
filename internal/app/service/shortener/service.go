package shortener

import (
	"database/sql"
	"net/url"

	"github.com/vabispklp/yap/internal/app/storage"
)

type Shortener struct {
	storage storage.StorageExpected

	db *sql.DB

	baseURL url.URL
}

func NewShortener(storage storage.StorageExpected, db *sql.DB, baseURL url.URL) (*Shortener, error) {
	if storage == nil {
		return nil, ErrNilPointerStorage
	}

	return &Shortener{
		storage: storage,
		db:      db,
		baseURL: baseURL,
	}, nil
}
