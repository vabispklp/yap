package handlers

type Handler struct {
	service ShortenerExpected
}

func NewHandler(service ShortenerExpected) (*Handler, error) {
	if service == nil {
		return nil, ErrNilPointerService
	}

	return &Handler{
		service: service,
	}, nil
}
