package model

type ShortURL struct {
	ID          string
	UserID      string
	OriginalURL string
}

type User struct {
	ID   string `json:"id"`
	Sign []byte `json:"key"`
}
