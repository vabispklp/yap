package shortener

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/vabispklp/yap/internal/app/model"
	storageMock "github.com/vabispklp/yap/internal/app/storage/mock"
)

func TestShortener_GetRedirectLink(t *testing.T) {
	type getStorageResult struct {
		shortURL *model.ShortURL
		err      error
	}
	type args struct {
		ctx context.Context
		id  string
	}

	var (
		id = "some_id"

		shortURL = &model.ShortURL{
			ID:          id,
			OriginalURL: "http://localhost:8080/some_path",
		}
	)

	tests := []struct {
		name             string
		getStorageResult getStorageResult
		args             args
		wantResult       *model.ShortURL
		wantErr          error
	}{
		{
			name: "успешное получение ссылки для редиректа",
			getStorageResult: getStorageResult{
				shortURL: nil,
				err:      nil,
			},
			args:       args{id: id},
			wantResult: nil,
			wantErr:    ErrNotFound,
		},
		{
			name: "storage вернул пустой объект",
			getStorageResult: getStorageResult{
				shortURL: shortURL,
				err:      nil,
			},
			args:       args{id: id},
			wantResult: shortURL,
			wantErr:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			storageMock := storageMock.NewMockStorageExpected(ctrl)

			storageMock.EXPECT().
				GetRedirectLink(gomock.Any(), gomock.Any()).
				Return(tt.getStorageResult.shortURL, tt.getStorageResult.err)

			s := &Shortener{
				storage: storageMock,
			}
			result, err := s.GetRedirectLink(tt.args.ctx, tt.args.id)

			assert.Equal(t, tt.wantResult, result, "Unexpected result")
			assert.Equal(t, tt.wantErr, err, "Unexpected error")

		})
	}
}

func TestShortener_AddRedirectLink(t *testing.T) {
	type args struct {
		ctx       context.Context
		stringURL string
	}
	tests := []struct {
		name             string
		addStorageResult error
		args             args
		expectedResult   string
		expectedErr      error
	}{
		{
			name:             "успешное сохранение короткой ссылки",
			addStorageResult: nil,
			args:             args{stringURL: "some_url"},
			expectedResult:   "",
			expectedErr:      nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			storageMock := storageMock.NewMockStorageExpected(ctrl)
			storageMock.EXPECT().
				AddRedirectLink(gomock.Any(), gomock.Any()).
				Return(tt.addStorageResult)

			s := &Shortener{
				storage: storageMock,
			}
			result, err := s.AddRedirectLink(tt.args.ctx, tt.args.stringURL)

			assert.NotEqual(t, nil, result, "Unexpected result")
			assert.Equal(t, tt.expectedErr, err, "Unexpected error")
		})
	}
}
