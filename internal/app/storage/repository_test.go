package storage

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/vabispklp/yap/internal/app/model"
)

func TestRepository_GetRedirectLink(t *testing.T) {
	var urlsMap = map[string]*model.ShortURL{
		"path": {
			Path:        "path",
			OriginalURL: "originalURL",
		},
	}
	type args struct {
		ctx  context.Context
		path string
	}
	tests := []struct {
		name           string
		args           args
		expectedResult *model.ShortURL
		expectedErr    error
	}{
		{
			name: "успешное получение короткой ссылки",
			args: args{path: "path"},
			expectedResult: &model.ShortURL{
				Path:        "path",
				OriginalURL: "originalURL",
			},
			expectedErr: nil,
		},
		{
			name:           "короткая ссылка на найдена",
			args:           args{path: "path 1"},
			expectedResult: nil,
			expectedErr:    errNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				urlsMap: urlsMap,
			}
			result, err := r.GetRedirectLink(tt.args.ctx, tt.args.path)

			assert.Equal(t, tt.expectedErr, err, "Unexpected error")
			assert.Equal(t, tt.expectedResult, result, "Unexpected result")
		})
	}
}

func TestRepository_AddRedirectLink(t *testing.T) {
	type fields struct {
		urlsMap map[string]*model.ShortURL
	}
	type args struct {
		ctx      context.Context
		shortURL *model.ShortURL
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		expectedErr error
	}{
		{
			name:   "успешное сохранение короткой ссылки",
			fields: fields{urlsMap: map[string]*model.ShortURL{}},
			args: args{
				shortURL: &model.ShortURL{
					Path:        "path",
					OriginalURL: "originalURL",
				},
			},
			expectedErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				urlsMap: tt.fields.urlsMap,
			}
			err := r.AddRedirectLink(tt.args.ctx, tt.args.shortURL)

			assert.Equal(t, tt.expectedErr, err, "Unexpected error")
		})
	}
}
