package handlers

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	shortenerMock "github.com/vabispklp/yap/cmd/shortener/handlers/mock"
	"github.com/vabispklp/yap/internal/app/model"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler_GetURL(t *testing.T) {
	type getServiceResult struct {
		shortURL *model.ShortURL
		err      error
	}
	type args struct {
		method string
		id     string
		url    string
	}
	type want struct {
		statusCode int
	}
	var tests = []struct {
		name             string
		getServiceResult getServiceResult
		args             args
		want             want
	}{
		{
			name: "успешный редирект по короткой ссылке",
			getServiceResult: getServiceResult{
				shortURL: &model.ShortURL{
					ID:          "some_id",
					OriginalURL: "https://google.com",
				},
				err: nil,
			},
			args: args{
				method: http.MethodGet,
				id:     "/some_id",
			},
			want: want{
				statusCode: http.StatusTemporaryRedirect,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			request := httptest.NewRequest(tt.args.method, tt.args.id, strings.NewReader(tt.args.url))

			w := httptest.NewRecorder()

			shortenerServiceMock := shortenerMock.NewMockShortenerExpected(ctrl)

			shortenerServiceMock.EXPECT().
				GetRedirectLink(gomock.Any(), gomock.Any()).
				Return(tt.getServiceResult.shortURL, tt.getServiceResult.err)

			h := Handler{service: shortenerServiceMock}

			h.GetHandleGetURL()(w, request)

			res := w.Result()
			defer res.Body.Close()

			assert.Equal(t, res.StatusCode, tt.want.statusCode, "Unexpected status code")
		})
	}
}
