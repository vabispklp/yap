package model

type (
	ShortenItemResponse struct {
		ShortURL    string `json:"short_url"`
		OriginalURL string `json:"original_url"`
	}

	ShortenBatchRequest struct {
		CorrelationID string `json:"correlation_id"`
		OriginalURL   string `json:"original_url"`
	}

	ShortenBatchResponse struct {
		CorrelationID string `json:"correlation_id"`
		ShortURL      string `json:"short_url"`
	}

	DeleteItemRequest struct {
		ID     string
		UserID string
	}
)
