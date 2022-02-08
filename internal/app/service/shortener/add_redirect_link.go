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

	shortURL, err := s.storage.Get(ctx, id)
	if err != nil {
		return "", err
	}
	if shortURL != nil && shortURL.UserID == userID {
		return u.String(), nil
	}

	err = s.storage.Add(ctx, model.ShortURL{
		ID:          id,
		UserID:      userID,
		OriginalURL: stringURL,
	})
	if err != nil {
		return "", err
	}

	return u.String(), nil
}
