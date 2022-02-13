package model

type ShortURL struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	OriginalURL string `json:"original_url"`
	Deleted     bool
}

type User struct {
	ID   string `json:"id"`
	Sign []byte `json:"key"`
}
