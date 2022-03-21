package handlers

import (
	"encoding/json"
	"github.com/vabispklp/yap/api/rest/middleware"
	"net/http"
)

// GetHandleGetUserURLs отдает хендлер который отдает все ссылки пользователя
func (h *Handler) GetHandleGetUserURLs() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		userID, ok := ctx.Value(middleware.ContextKeyUserID).(string)
		if !ok {
			http.Error(w, errTextInternal, http.StatusInternalServerError)
			return
		}

		shortURLs, err := h.service.GetUserURLs(ctx, userID)
		if err != nil {
			http.Error(w, errTextInternal, http.StatusInternalServerError)
			return
		}

		if shortURLs == nil {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		res, err := json.Marshal(shortURLs)
		if err != nil {
			http.Error(w, errTextInternal, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	}
}
