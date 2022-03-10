package handlers

import (
	"encoding/json"
	"github.com/vabispklp/yap/api/rest/middleware"
	"github.com/vabispklp/yap/internal/app/service/model"
	"io"
	"net/http"
	"net/url"
)

// GetHandlerAddBatch отдает хендлер который занимается можественным добавлением сокращенных ссылок
func (h *Handler) GetHandlerAddBatch() func(w http.ResponseWriter, r *http.Request) {
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

		var requestItems []model.ShortenBatchRequest
		if err = json.Unmarshal(b, &requestItems); err != nil {
			http.Error(w, errTextEmptyURL, http.StatusBadRequest)
			return
		}

		for _, item := range requestItems {
			if _, err = url.Parse(item.OriginalURL); err != nil {
				http.Error(w, errTextInvalidOriginalURL, http.StatusBadRequest)
				return
			}

			if item.CorrelationID == "" {
				http.Error(w, errTextInvalidCorrelationID, http.StatusBadRequest)
				return
			}
		}

		userID, ok := ctx.Value(middleware.ContextKeyUserID).(string)
		if !ok {
			http.Error(w, errTextInternal, http.StatusInternalServerError)
			return
		}

		result, err := h.service.AddManyRedirectLink(ctx, requestItems, userID)
		if err != nil {
			http.Error(w, errTextInternal, http.StatusInternalServerError)
			return
		}

		b, err = json.Marshal(result)
		if err != nil {
			http.Error(w, errTextInternal, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(b)
	}
}
