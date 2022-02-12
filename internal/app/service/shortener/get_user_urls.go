package shortener

import (
	"context"
	"github.com/vabispklp/yap/internal/app/service/model"
)

func (s *Shortener) GetUserURLs(ctx context.Context, userID string) ([]model.Shorten, error) {
	shortURLs, err := s.storage.GetByUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	if len(shortURLs) == 0 {
		return nil, nil
	}

	u := s.baseURL

	result := make([]model.Shorten, 0, len(shortURLs))
	for _, item := range shortURLs {
		u.Path = item.ID
		result = append(result, model.Shorten{
			ShortURL:    u.String(),
			OriginalURL: item.OriginalURL,
		})
	}

	return result, nil
}
