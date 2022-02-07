package shortener

import (
	"context"
	"crypto/md5"
	"encoding/hex"

	"github.com/vabispklp/yap/internal/app/model"
)

func (s *Shortener) AddRedirectLink(ctx context.Context, stringURL string) (string, error) {
	hash := md5.Sum([]byte(stringURL))
	resultPath := hex.EncodeToString(hash[:])

	u := s.baseURL
	u.Path = resultPath

	err := s.storage.Add(ctx, model.ShortURL{
		ID:          resultPath,
		OriginalURL: stringURL,
	})
	if err != nil {
		return "", err
	}

	return u.String(), nil
}
