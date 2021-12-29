package handlers

import "errors"

const (
	errTextInternal   = "Internal error"
	errTextEmptyID    = "Empty id"
	errTextInvalidURL = `Parameter "url" is invalid`
	errTextEmptyURL   = `Parameter "url" is required`
	errTextEmptyBody  = "Empty body"
)

var ErrNilPointerService = errors.New("nil pointer service")
