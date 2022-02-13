package shortener

import (
	"context"
	"github.com/vabispklp/yap/internal/app/storage/model"
)

func (s *Shortener) GetRedirectLink(ctx context.Context, id string) (*model.ShortURL, error) {
	result, err := s.storage.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, ErrNotFound
	}

	if result.Deleted {
		return nil, ErrDeleted
	}

	return result, nil
}
