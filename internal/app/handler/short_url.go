package handler

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/vabispklp/yap/internal/app/model"
)

const (
	resultURLPattern = "http://localhost:8080/%s"
)

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	switch r.Method {
	case http.MethodPost:
		if r.Body == nil {
			http.Error(w, "Empty body", http.StatusBadRequest)

			return
		}

		defer r.Body.Close()

		b, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Internal error", http.StatusInternalServerError)

			return
		}

		stringURL := string(b)
		if stringURL == "" {
			http.Error(w, `Parameter "url" is required`, http.StatusBadRequest)

			return
		}

		_, err = url.Parse(stringURL)
		if err != nil {
			http.Error(w, `Parameter "url" is invalid`, http.StatusBadRequest)

			return
		}

		hash := md5.Sum([]byte(stringURL))
		resultPath := hex.EncodeToString(hash[:])

		resultURL := fmt.Sprintf(resultURLPattern, resultPath)

		err = h.storage.AddRedirectLink(ctx, &model.ShortURL{
			Path:        resultPath,
			OriginalURL: stringURL,
		})
		if err != nil {
			http.Error(w, `Internal error`, http.StatusInternalServerError)

			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(resultURL))
	case http.MethodGet:
		requestURI := r.URL.RequestURI()
		if requestURI == "" {
			http.Error(w, `Empty path`, http.StatusBadRequest)

			return
		}

		shortURL, err := h.storage.GetRedirectLink(ctx, strings.TrimLeft(requestURI, "/"))
		if err != nil {
			http.NotFound(w, r)

			return
		}

		http.Redirect(w, r, shortURL.OriginalURL, http.StatusTemporaryRedirect)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
