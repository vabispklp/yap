package shortener

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	storageMock "github.com/vabispklp/yap/internal/app/storage/mock"
)

func TestShortener_Ping(t *testing.T) {
	type getStorageResult struct {
		err error
	}
	type args struct {
		ctx context.Context
	}

	tests := []struct {
		name             string
		getStorageResult getStorageResult
		args             args
		wantErr          error
	}{
		{
			name: "успешный пинг",
			getStorageResult: getStorageResult{
				err: nil,
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			storageMock := storageMock.NewMockStorageExpected(ctrl)

			storageMock.EXPECT().
				Ping(gomock.Any()).
				Return(tt.getStorageResult.err)

			s := &Shortener{
				storage: storageMock,
			}
			err := s.Ping(tt.args.ctx)

			assert.Equal(t, tt.wantErr, err, "Unexpected error")

		})
	}
}
