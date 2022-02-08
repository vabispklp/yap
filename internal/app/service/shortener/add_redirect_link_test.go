package shortener

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	storageMock "github.com/vabispklp/yap/internal/app/storage/mock"
)

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
				Get(gomock.Any(), gomock.Any()).
				Return(nil, nil)

			storageMock.EXPECT().
				Add(gomock.Any(), gomock.Any()).
				Return(tt.addStorageResult)

			s := &Shortener{
				storage: storageMock,
			}

			result, err := s.AddRedirectLink(tt.args.ctx, tt.args.stringURL, "")

			assert.NotEqual(t, nil, result, "Unexpected result")
			assert.Equal(t, tt.expectedErr, err, "Unexpected error")
		})
	}
}
