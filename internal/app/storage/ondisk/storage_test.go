package ondisk

import (
	"context"
	"github.com/vabispklp/yap/internal/app/storage/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStorage_Add(t *testing.T) {
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
					ID:          "some_id_1",
					OriginalURL: "originalURL1",
				},
			},
			expectedErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, _ := NewStorage("/tmp/test")
			err := r.Add(tt.args.ctx, tt.args.shortURL)
			defer r.Close()

			assert.Equal(t, tt.expectedErr, err, "Unexpected error")
		})
	}
}

func TestStorage_Get(t *testing.T) {
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
			args: args{id: "some_id_2"},
			expectedResult: &model.ShortURL{
				ID:          "some_id_2",
				OriginalURL: "originalURL2",
			},
			expectedErr: nil,
		},
		{
			name:           "короткая ссылка на найдена",
			args:           args{id: "some_id_100"},
			expectedResult: nil,
			expectedErr:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, _ := NewStorage("/tmp/test")
			r.Add(tt.args.ctx, model.ShortURL{
				ID:          "some_id_2",
				OriginalURL: "originalURL2",
			})
			result, err := r.Get(tt.args.ctx, tt.args.id)

			defer r.Close()

			assert.Equal(t, tt.expectedErr, err, "Unexpected error")
			assert.Equal(t, tt.expectedResult, result, "Unexpected result")
		})
	}
}

func TestStorage_GetByUser(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID string
	}
	tests := []struct {
		name           string
		args           args
		expectedResult []model.ShortURL
		expectedErr    error
	}{
		{
			name: "успешное получение коротих ссылок пользователя",
			args: args{userID: "some_user_id"},
			expectedResult: []model.ShortURL{
				{
					ID:          "some_id_3",
					OriginalURL: "originalURL3",
					UserID:      "some_user_id",
				},
			},
			expectedErr: nil,
		},
		{
			name:           "короткие ссылки на найдены",
			args:           args{userID: "some_id_100"},
			expectedResult: []model.ShortURL{},
			expectedErr:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, _ := NewStorage("/tmp/test")
			r.Add(tt.args.ctx, model.ShortURL{
				ID:          "some_id_3",
				OriginalURL: "originalURL3",
				UserID:      "some_user_id",
			})
			result, err := r.GetByUser(tt.args.ctx, tt.args.userID)

			defer r.Close()

			assert.Equal(t, tt.expectedErr, err, "Unexpected error")
			assert.Equal(t, tt.expectedResult, result, "Unexpected result")
		})
	}
}
