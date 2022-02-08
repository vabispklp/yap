package inmem

import (
	"context"
	"github.com/vabispklp/yap/internal/app/storage/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStorage_GetRedirectLink(t *testing.T) {
	var urlsMap = map[string]model.ShortURL{
		"some_id": {
			ID:          "some_id",
			OriginalURL: "originalURL",
		},
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name           string
		args           args
		expectedResult *model.ShortURL
		expectedErr    error
	}{
		{
			name: "успешное получение короткой ссылки",
			args: args{id: "some_id"},
			expectedResult: &model.ShortURL{
				ID:          "some_id",
				OriginalURL: "originalURL",
			},
			expectedErr: nil,
		},
		{
			name:           "короткая ссылка на найдена",
			args:           args{id: "some_id_2"},
			expectedResult: nil,
			expectedErr:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Storage{
				urlsMap: urlsMap,
			}
			result, err := r.Get(tt.args.ctx, tt.args.id)

			assert.Equal(t, tt.expectedErr, err, "Unexpected error")
			assert.Equal(t, tt.expectedResult, result, "Unexpected result")
		})
	}
}

func TestStorage_AddRedirectLink(t *testing.T) {
	type fields struct {
		urlsMap map[string]model.ShortURL
	}
	type args struct {
		ctx      context.Context
		shortURL model.ShortURL
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		expectedErr error
	}{
		{
			name:   "успешное сохранение короткой ссылки",
			fields: fields{urlsMap: map[string]model.ShortURL{}},
			args: args{
				shortURL: model.ShortURL{
					ID:          "some_id",
					OriginalURL: "originalURL",
				},
			},
			expectedErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Storage{
				urlsMap: tt.fields.urlsMap,
			}
			err := r.Add(tt.args.ctx, tt.args.shortURL)

			assert.Equal(t, tt.expectedErr, err, "Unexpected error")
		})
	}
}
