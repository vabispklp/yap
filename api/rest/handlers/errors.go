package handlers

import "errors"

const (
	errTextInternal             = "Internal error"
	errTextEmptyID              = "Empty id"
	errTextInvalidURL           = `Parameter "url" is invalid`
	errTextInvalidOriginalURL   = `Parameter "original_url" is invalid`
	errTextInvalidCorrelationID = `Parameter "correlation_id" is invalid`
	errTextEmptyURL             = `Parameter "url" is required`
	errTextEmptyBody            = "Empty body"
	errInvalidRequestBody       = "Invalid request body"
)

var ErrNilPointerService = errors.New("nil pointer service")
