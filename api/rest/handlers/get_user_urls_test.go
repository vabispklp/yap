package handlers

import (
	"context"
	"github.com/go-chi/chi/v5"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	shortenerMock "github.com/vabispklp/yap/api/rest/handlers/mock"
	"github.com/vabispklp/yap/api/rest/middleware"
	"github.com/vabispklp/yap/internal/app/service/model"
)

func ExampleGetHandleGetUserURLs() {
	// Создаем любой роутер
	router := chi.NewRouter()

	// Создаем струтуру хендлеров
	h, _ := NewHandler(ShortenerExpected(nil))

	// Получаем хендлер всех ссылок пользователя
	router.Get("/some_route", h.GetHandleGetUserURLs())
}

func TestHandler_GetUserURLs(t *testing.T) {
	type getUserURLsResult struct {
		urls []model.ShortenItemResponse
		err  error
	}
	type args struct {
		method string
		target string
	}
	type want struct {
		statusCode int
		response   string
	}
	var tests = []struct {
		name              string
		getUserURLsResult getUserURLsResult
		args              args
		want              want
	}{
		{
			name: "успешное получение ссылок пользователя",
			getUserURLsResult: getUserURLsResult{
				urls: []model.ShortenItemResponse{
					{
						OriginalURL: "http://localhost:8080/some",
						ShortURL:    "http://localhost:8080/short_some",
					},
				},
				err: nil,
			},
			args: args{
				method: http.MethodGet,
				target: "/user/urls",
			},
			want: want{
				statusCode: http.StatusOK,
				response:   `[{"short_url":"http://localhost:8080/short_some","original_url":"http://localhost:8080/some"}]`,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			request := httptest.NewRequest(tt.args.method, tt.args.target, nil)

			w := httptest.NewRecorder()

			shortenerServiceMock := shortenerMock.NewMockShortenerExpected(ctrl)

			shortenerServiceMock.EXPECT().
				GetUserURLs(gomock.Any(), gomock.Any()).
				Return(tt.getUserURLsResult.urls, tt.getUserURLsResult.err)

			h := Handler{service: shortenerServiceMock}
			ctx := request.Context()

			h.GetHandleGetUserURLs()(w, request.WithContext(context.WithValue(ctx, middleware.ContextKeyUserID, "someUserID")))

			res := w.Result()
			defer res.Body.Close()
			result, err := ioutil.ReadAll(res.Body)

			require.NoError(t, err, "decode has error")
			assert.Equal(t, res.StatusCode, tt.want.statusCode, "Unexpected status code")
			assert.NotEqual(t, string(result), "", "Unexpected result")
		})
	}
}
