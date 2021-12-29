package handlers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/vabispklp/yap/internal/app/service/shortener"
)

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
			} else {
				http.Error(w, errTextInternal, http.StatusInternalServerError)
			}

			return
		}

		http.Redirect(w, r, shortURL.OriginalURL, http.StatusTemporaryRedirect)
	}
}
