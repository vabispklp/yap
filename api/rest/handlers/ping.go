package handlers

import (
	"net/http"
)

func (h *Handler) GetHandlerPing() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h.service.Ping(r.Context()); err != nil {
			http.Error(w, errTextInternal, http.StatusInternalServerError)
			return
		}
	}
}
