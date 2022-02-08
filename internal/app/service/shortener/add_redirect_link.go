package shortener

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"github.com/vabispklp/yap/internal/app/storage/model"
)

func (s *Shortener) AddRedirectLink(ctx context.Context, stringURL, userID string) (string, error) {
	hash := md5.Sum([]byte(stringURL))
	id := hex.EncodeToString(hash[:])

	u := s.baseURL
	u.Path = id

	err := s.storage.Add(ctx, model.ShortURL{
		ID:          id,
		UserID:      userID,
		OriginalURL: stringURL,
	})

	if err != nil {
		return "", err
	}

	return u.String(), nil
}
