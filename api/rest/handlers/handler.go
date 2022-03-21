package handlers

// Handler структура с коллекцией хендлеров
type Handler struct {
	service ShortenerExpected
}

// NewHandler создает Handler с коллекцией хендлеров
func NewHandler(service ShortenerExpected) (*Handler, error) {
	if service == nil {
		return nil, ErrNilPointerService
	}

	return &Handler{
		service: service,
	}, nil
}
