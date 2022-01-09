package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	shortenerMock "github.com/vabispklp/yap/api/rest/handlers/mock"
)

func TestHandler_AddShorten(t *testing.T) {
	type addServiceResult struct {
		url string
		err error
	}
	type args struct {
		method  string
		target  string
		request AddShortenRequest
	}
	type want struct {
		statusCode int
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
				url: "http://localhost:8080/short",
				err: nil,
			},
			args: args{
				method:  http.MethodPost,
				target:  "/",
				request: AddShortenRequest{URL: "http://localhost:8080/some_id"},
			},
			want: want{
				statusCode: http.StatusCreated,
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
				AddRedirectLink(gomock.Any(), gomock.Any()).
				Return(tt.addStorageResult.url, tt.addStorageResult.err)

			h := Handler{service: shortenerServiceMock}

			h.GetHandlerAddShorten()(w, request)

			res := w.Result()
			defer res.Body.Close()
			result, err := ioutil.ReadAll(res.Body)

			require.Nil(t, err, "decode error not nil")
			assert.Equal(t, res.StatusCode, tt.want.statusCode, "Unexpected status code")
			assert.NotEqual(t, string(result), "", "Unexpected result")
		})
	}
}
