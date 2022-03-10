package handlers

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	shortenerMock "github.com/vabispklp/yap/api/rest/handlers/mock"
)

func ExampleGetHandlerPing() {
	// Создаем любой роутер
	router := chi.NewRouter()

	// Создаем струтуру хендлеров
	h, _ := NewHandler(ShortenerExpected(nil))

	// Получаем хендлер для пинга
	router.Post("/some_route", h.GetHandlerPing())
}

func TestHandler_Ping(t *testing.T) {
	type pingResult struct {
		err error
	}
	type args struct {
		method string
		target string
	}
	type want struct {
		statusCode int
	}
	var tests = []struct {
		name       string
		pingResult pingResult
		args       args
		want       want
	}{
		{
			name: "успешный пинг",
			pingResult: pingResult{
				err: nil,
			},
			args: args{
				method: http.MethodGet,
				target: "/ping",
			},

			want: want{
				statusCode: http.StatusOK,
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
				Ping(gomock.Any()).
				Return(tt.pingResult.err)

			h := Handler{service: shortenerServiceMock}

			h.GetHandlerPing()(w, request)

			res := w.Result()
			defer res.Body.Close()

			assert.Equal(t, res.StatusCode, tt.want.statusCode, "Unexpected status code")
		})
	}
}
