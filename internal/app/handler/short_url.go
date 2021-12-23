package handler

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

const (
	resultURLPattern = "http://localhost:8080/%s"
)

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

		mu := sync.Mutex{}
		mu.Lock()
		h.urlsMap[resultPath] = stringURL
		mu.Unlock()

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(resultURL))
	case http.MethodGet:
		requestURI := r.URL.RequestURI()
		if requestURI == "" {
			http.Error(w, `Empty path`, http.StatusBadRequest)

			return
		}

		a := strings.TrimLeft(requestURI, "/")
		originalURL, ok := h.urlsMap[a]
		if !ok {
			http.NotFound(w, r)

			return
		}

		http.Redirect(w, r, originalURL, http.StatusTemporaryRedirect)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
