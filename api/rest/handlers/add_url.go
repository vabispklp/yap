package handlers

import (
	"io"
	"net/http"
	"net/url"
)

func (h *Handler) GetHandlerAddURL() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		if r.Body == nil {
			http.Error(w, errTextEmptyBody, http.StatusBadRequest)
			return
		}

		b, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, errTextInternal, http.StatusInternalServerError)
			return
		}

		defer r.Body.Close()

		stringURL := string(b)
		if stringURL == "" {
			http.Error(w, errTextEmptyURL, http.StatusBadRequest)
			return
		}

		if _, err = url.Parse(stringURL); err != nil {
			http.Error(w, errTextInvalidURL, http.StatusBadRequest)
			return
		}

		resultURL, err := h.service.AddRedirectLink(ctx, stringURL)
		if err != nil {
			http.Error(w, errTextInternal, http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(resultURL))
	}
}
