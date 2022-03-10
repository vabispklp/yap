package shortener

import (
	"context"
	"github.com/vabispklp/yap/internal/app/storage/model"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	storageMock "github.com/vabispklp/yap/internal/app/storage/mock"
)

func TestShortener_DeleteRedirectLink(t *testing.T) {
	type getStorageResult struct {
		url *model.ShortURL
		err error
	}

	type args struct {
		ctx context.Context
		ids []string
	}
	tests := []struct {
		name                string
		getStorageResult    getStorageResult
		deleteStorageResult error
		args                args
		expectedResult      string
		expectedErr         error
	}{
		{
			name:                "успешное удаление коротих ссылок",
			deleteStorageResult: nil,
			getStorageResult: getStorageResult{
				url: &model.ShortURL{
					ID:          "some_id",
					OriginalURL: "some_url",
				},
				err: nil,
			},
			args: args{
				ids: []string{"some_id"},
			},
			expectedResult: "",
			expectedErr:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			storageMock := storageMock.NewMockStorageExpected(ctrl)
			storageMock.EXPECT().
				Get(gomock.Any(), gomock.Any()).
				Return(tt.getStorageResult.url, tt.getStorageResult.err)

			storageMock.EXPECT().
				Delete(gomock.Any(), gomock.Any(), gomock.Any()).
				Return(tt.deleteStorageResult)

			s := &Shortener{
				storage: storageMock,
			}

			err := s.DeleteRedirectLinks(tt.args.ctx, tt.args.ids, "")

			assert.Equal(t, tt.expectedErr, err, "Unexpected error")
		})
	}
}
