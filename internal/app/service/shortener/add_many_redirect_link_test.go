package shortener

import (
	"context"
	"github.com/vabispklp/yap/internal/app/service/model"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	storageMock "github.com/vabispklp/yap/internal/app/storage/mock"
)

func TestShortener_AddManyRedirectLink(t *testing.T) {
	type args struct {
		ctx  context.Context
		urls []model.ShortenBatchRequest
	}
	tests := []struct {
		name             string
		addStorageResult error
		args             args
		expectedResult   string
		expectedErr      error
	}{
		{
			name:             "успешное сохранение коротих ссылок",
			addStorageResult: nil,
			args: args{
				urls: []model.ShortenBatchRequest{
					{
						CorrelationID: "some_correlation_id",
						OriginalURL:   "some_url",
					},
				},
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
				AddMany(gomock.Any(), gomock.Any()).
				Return(tt.addStorageResult)

			s := &Shortener{
				storage: storageMock,
			}

			result, err := s.AddManyRedirectLink(tt.args.ctx, tt.args.urls, "")

			assert.NotEqual(t, nil, result, "Unexpected result")
			assert.Equal(t, tt.expectedErr, err, "Unexpected error")
		})
	}
}
