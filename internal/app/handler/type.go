package handler

type Handler struct {
	urlsMap map[string]string
}

func New() *Handler {
	return &Handler{
		urlsMap: make(map[string]string),
	}
}
