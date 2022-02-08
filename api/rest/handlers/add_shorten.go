package handlers

import (
	"encoding/json"
	"github.com/jackc/pgerrcode"
	"github.com/lib/pq"
	"github.com/vabispklp/yap/api/rest/middleware"
	"io"
	"net/http"
	"net/url"
)

type AddShortenRequest struct {
	URL string `json:"url"`
}

type AddShortenResponse struct {
	Result string `json:"result"`
}

func (h *Handler) GetHandlerAddShorten() func(w http.ResponseWriter, r *http.Request) {
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

		var args AddShortenRequest
		if err = json.Unmarshal(b, &args); err != nil {
			http.Error(w, errTextEmptyURL, http.StatusBadRequest)
			return
		}

		if _, err = url.Parse(args.URL); err != nil {
			http.Error(w, errTextInvalidURL, http.StatusBadRequest)
			return
		}

		userID, ok := ctx.Value(middleware.ContextKeyUserID).(string)
		if !ok {
			http.Error(w, errTextInternal, http.StatusInternalServerError)
			return
		}

		status := http.StatusCreated

		resultURL, err := h.service.AddRedirectLink(ctx, args.URL, userID)
		if err != nil {
			// всегда false
			//if errors.As(err, &pq.Error{}) {}
			pqErr, ok := err.(*pq.Error)
			if ok && pqErr.Code == pgerrcode.UniqueViolation {
				status = http.StatusConflict
			} else {
				http.Error(w, errTextInternal, http.StatusInternalServerError)
				return
			}
		}

		b, err = json.Marshal(AddShortenResponse{
			Result: resultURL,
		})
		if err != nil {
			http.Error(w, errTextInternal, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		w.Write(b)
	}
}
