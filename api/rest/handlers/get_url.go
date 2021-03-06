package handlers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/vabispklp/yap/internal/app/service/shortener"
)

// GetHandleGetURL отдает хендлер который занимается получением оригинальных ссылок и редиректом на них
func (h *Handler) GetHandleGetURL() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		id := r.URL.RequestURI()
		if id == "" {
			http.Error(w, errTextEmptyID, http.StatusBadRequest)
			return
		}

		shortURL, err := h.service.GetRedirectLink(ctx, strings.TrimLeft(id, "/"))
		if err != nil {
			if errors.Is(err, shortener.ErrNotFound) {
				http.NotFound(w, r)
			} else if errors.Is(err, shortener.ErrDeleted) {
				w.WriteHeader(http.StatusGone)
				return
			} else {
				http.Error(w, errTextInternal, http.StatusInternalServerError)
			}

			return
		}

		http.Redirect(w, r, shortURL.OriginalURL, http.StatusTemporaryRedirect)
	}
}
