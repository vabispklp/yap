package handlers

import (
	"encoding/json"
	"github.com/vabispklp/yap/api/rest/middleware"
	"io"
	"net/http"
)

// GetHandlerDelete отдает хендлер который занимается удалением ссылок
func (h *Handler) GetHandlerDelete() func(w http.ResponseWriter, r *http.Request) {
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

		var requestIDs []string
		if err = json.Unmarshal(b, &requestIDs); err != nil {
			http.Error(w, errInvalidRequestBody, http.StatusBadRequest)
			return
		}

		userID, ok := ctx.Value(middleware.ContextKeyUserID).(string)
		if !ok {
			http.Error(w, errTextInternal, http.StatusInternalServerError)
			return
		}

		err = h.service.DeleteRedirectLinks(ctx, requestIDs, userID)
		if err != nil {
			http.Error(w, errTextInternal, http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusAccepted)
	}
}
