package handlers

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	shortenerMock "github.com/vabispklp/yap/api/rest/handlers/mock"
	"github.com/vabispklp/yap/api/rest/middleware"
	"github.com/vabispklp/yap/internal/app/service/model"
)

func ExampleHandler_GetHandlerAddBatch() {
	// Создаем любой роутер
	router := chi.NewRouter()

	// Создаем струтуру хендлеров
	h, _ := NewHandler(ShortenerExpected(nil))

	// Получаем хендлер множественного сокращения ссылок
	router.Post("/some_route", h.GetHandlerAddBatch())
}

func TestHandler_AddBatch(t *testing.T) {
	type addServiceResult struct {
		urls []model.ShortenBatchResponse
		err  error
	}
	type args struct {
		method  string
		target  string
		request []model.ShortenBatchRequest
	}
	type want struct {
		statusCode int
		response   string
	}
	var tests = []struct {
		name             string
		addStorageResult addServiceResult
		args             args
		want             want
	}{
		{
			name: "успешное добавление короткой ссылки",
			addStorageResult: addServiceResult{
				urls: []model.ShortenBatchResponse{
					{
						CorrelationID: "123",
						ShortURL:      "http://localhost:8080/short_some",
					},
				},
				err: nil,
			},
			args: args{
				method: http.MethodPost,
				target: "/api/shorten/batch",
				request: []model.ShortenBatchRequest{
					{
						CorrelationID: "123", OriginalURL: "http://localhost:8080/some",
					},
				},
			},
			want: want{
				statusCode: http.StatusCreated,
				response:   `[{"correlation_id":"123","short_url":"http://localhost:8080/short_some"}]`,
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
				AddManyRedirectLink(gomock.Any(), gomock.Any(), gomock.Any()).
				Return(tt.addStorageResult.urls, tt.addStorageResult.err)

			h := Handler{service: shortenerServiceMock}
			ctx := request.Context()

			h.GetHandlerAddBatch()(w, request.WithContext(context.WithValue(ctx, middleware.ContextKeyUserID, "someUserID")))

			res := w.Result()
			defer res.Body.Close()
			result, err := ioutil.ReadAll(res.Body)

			require.NoError(t, err, "decode has error")
			assert.Equal(t, res.StatusCode, tt.want.statusCode, "Unexpected status code")
			assert.NotEqual(t, string(result), "", "Unexpected result")
		})
	}
}
