package shortener

import "errors"

var (
	ErrNilPointerStorage = errors.New("nil pointer storage")

	ErrNotFound = errors.New("original URL not found")
)
