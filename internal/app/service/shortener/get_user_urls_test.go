package shortener

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	serviceModel "github.com/vabispklp/yap/internal/app/service/model"
	storageMock "github.com/vabispklp/yap/internal/app/storage/mock"
	"github.com/vabispklp/yap/internal/app/storage/model"
)

func TestShortener_GetUserUrls(t *testing.T) {
	type getStorageResult struct {
		shortURLs []model.ShortURL
		err       error
	}
	type args struct {
		ctx    context.Context
		userID string
	}

	var (
		userID = "some_user_id"

		shortURL = model.ShortURL{
			ID:          userID,
			OriginalURL: "http://localhost:8080/some_path",
			UserID:      userID,
		}
	)

	tests := []struct {
		name             string
		getStorageResult getStorageResult
		args             args
		wantResult       []serviceModel.ShortenItemResponse
		wantErr          error
	}{
		{
			name: "успешное получение ссылок пользователей",
			getStorageResult: getStorageResult{
				shortURLs: []model.ShortURL{shortURL},
				err:       nil,
			},
			args: args{userID: userID},
			wantResult: []serviceModel.ShortenItemResponse{{
				ShortURL:    shortURL.ID,
				OriginalURL: shortURL.OriginalURL,
			}},
			wantErr: nil,
		},
		{
			name: "storage вернул пустой объект",
			getStorageResult: getStorageResult{
				shortURLs: nil,
				err:       nil,
			},
			args:       args{userID: userID},
			wantResult: nil,
			wantErr:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			storageMock := storageMock.NewMockStorageExpected(ctrl)

			storageMock.EXPECT().
				GetByUser(gomock.Any(), gomock.Any()).
				Return(tt.getStorageResult.shortURLs, tt.getStorageResult.err)

			s := &Shortener{
				storage: storageMock,
			}
			result, err := s.GetUserURLs(tt.args.ctx, tt.args.userID)

			assert.Equal(t, tt.wantResult, result, "Unexpected result")
			assert.Equal(t, tt.wantErr, err, "Unexpected error")

		})
	}
}
