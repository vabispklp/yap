package shortener

import (
	"context"
	"crypto/md5"
	"encoding/hex"

	"github.com/vabispklp/yap/internal/app/service/model"
	storageModel "github.com/vabispklp/yap/internal/app/storage/model"
)

func (s *Shortener) AddManyRedirectLink(ctx context.Context, shortenBatchItems []model.ShortenBatch, userID string) ([]model.ShortenBatchResult, error) {
	shortURLs := make([]storageModel.ShortURL, 0, len(shortenBatchItems))
	result := make([]model.ShortenBatchResult, 0, len(shortenBatchItems))
	for _, item := range shortenBatchItems {
		hash := md5.Sum([]byte(item.OriginalURL))
		id := hex.EncodeToString(hash[:])

		u := s.baseURL
		u.Path = id

		shortURLs = append(shortURLs, storageModel.ShortURL{
			ID:          id,
			UserID:      userID,
			OriginalURL: item.OriginalURL,
		})

		result = append(result, model.ShortenBatchResult{
			CorrelationID: item.CorrelationID,
			ShortURL:      u.String(),
		})
	}

	err := s.storage.AddMany(ctx, shortURLs)
	if err != nil {
		return nil, err
	}

	return result, nil
}
