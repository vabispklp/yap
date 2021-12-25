package handler

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vabispklp/yap/internal/app/model"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler_ServeHTTPPost(t *testing.T) {
	type args struct {
		method string
		path   string
		url    string
	}
	type wont struct {
		statusCode int
	}
	var tests = []struct {
		name             string
		addStorageResult error
		args             args
		wont             wont
	}{
		{
			name:             "успешное добавление короткой ссылки",
			addStorageResult: nil,
			args: args{
				method: http.MethodPost,
				path:   "/",
				url:    "http://localhost:8080/abc",
			},
			wont: wont{
				statusCode: http.StatusCreated,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(tt.args.method, tt.args.path, strings.NewReader(tt.args.url))

			w := httptest.NewRecorder()

			h := Handler{storage: &MockStorage{
				AddFunc: func(ctx context.Context, url *model.ShortURL) error {
					return nil
				},
			}}

			h.ServeHTTP(w, request)

			res := w.Result()
			defer res.Body.Close()
			result, err := ioutil.ReadAll(res.Body)

			require.Nil(t, err, "decode error not nil")
			assert.Equal(t, res.StatusCode, tt.wont.statusCode, "Unexpected status code")
			assert.NotEqual(t, string(result), "", "Unexpected result")
		})
	}
}

func TestHandler_ServeHTTPGet(t *testing.T) {
	type getStorageResult struct {
		shortURL *model.ShortURL
		err      error
	}
	type args struct {
		method string
		path   string
		url    string
	}
	type wont struct {
		statusCode int
	}
	var tests = []struct {
		name             string
		getStorageResult getStorageResult
		args             args
		wont             wont
	}{
		{
			name: "успешный редирект по короткой ссылке",
			getStorageResult: getStorageResult{
				shortURL: &model.ShortURL{
					Path:        "path",
					OriginalURL: "https://google.com",
				},
				err: nil,
			},
			args: args{
				method: http.MethodGet,
				path:   "/path",
			},
			wont: wont{
				statusCode: http.StatusTemporaryRedirect,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(tt.args.method, tt.args.path, strings.NewReader(tt.args.url))

			w := httptest.NewRecorder()

			h := Handler{storage: &MockStorage{

				GetFunc: func(ctx context.Context, s string) (*model.ShortURL, error) {
					return tt.getStorageResult.shortURL, tt.getStorageResult.err
				},
			}}

			h.ServeHTTP(w, request)

			res := w.Result()
			defer res.Body.Close()

			assert.Equal(t, res.StatusCode, tt.wont.statusCode, "Unexpected status code")
		})
	}
}

type MockStorage struct {
	GetFunc func(context.Context, string) (*model.ShortURL, error)
	AddFunc func(context.Context, *model.ShortURL) error
}

// GetRedirectLink calls m.GetRedirectLink.
func (m *MockStorage) GetRedirectLink(ctx context.Context, path string) (*model.ShortURL, error) {
	return m.GetFunc(ctx, path)
}

// AddRedirectLink calls m.AddRedirectLink.
func (m *MockStorage) AddRedirectLink(ctx context.Context, shortURL *model.ShortURL) error {
	return m.AddFunc(ctx, shortURL)
}
