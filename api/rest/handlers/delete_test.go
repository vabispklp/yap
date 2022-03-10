package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	shortenerMock "github.com/vabispklp/yap/api/rest/handlers/mock"
	"github.com/vabispklp/yap/api/rest/middleware"
)

func ExampleHandler_GetHandlerDelete() {
	// Создаем любой роутер
	router := chi.NewRouter()

	// Создаем струтуру хендлеров
	h, _ := NewHandler(ShortenerExpected(nil))

	// Получаем хендлер удаления ссылки
	router.Delete("/some_route", h.GetHandlerDelete())
}

func TestHandler_Delete(t *testing.T) {
	type deleteResult struct {
		err error
	}
	type args struct {
		method  string
		target  string
		request []string
	}
	type want struct {
		statusCode int
	}
	var tests = []struct {
		name         string
		deleteResult deleteResult
		args         args
		want         want
	}{
		{
			name: "успешное удаление короткой ссылки",
			deleteResult: deleteResult{
				err: nil,
			},
			args: args{
				method:  http.MethodDelete,
				target:  "/api/user/urls",
				request: []string{"123"},
			},

			want: want{
				statusCode: http.StatusAccepted,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			requestBody, _ := json.Marshal(tt.args.request)
			request := httptest.NewRequest(tt.args.method, tt.args.target, strings.NewReader(string(requestBody)))

			w := httptest.NewRecorder()

			shortenerServiceMock := shortenerMock.NewMockShortenerExpected(ctrl)

			shortenerServiceMock.EXPECT().
				DeleteRedirectLinks(gomock.Any(), gomock.Any(), gomock.Any()).
				Return(tt.deleteResult.err)

			h := Handler{service: shortenerServiceMock}
			ctx := request.Context()

			h.GetHandlerDelete()(w, request.WithContext(context.WithValue(ctx, middleware.ContextKeyUserID, "someUserID")))

			res := w.Result()
			defer res.Body.Close()

			assert.Equal(t, res.StatusCode, tt.want.statusCode, "Unexpected status code")
		})
	}
}
