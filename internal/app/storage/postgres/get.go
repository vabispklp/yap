package postgres

import (
	"context"
	"github.com/vabispklp/yap/internal/app/storage/model"
)

// Get отдает сокращенную ссылку
func (s *Storage) Get(ctx context.Context, id string) (*model.ShortURL, error) {
	row := s.db.QueryRowContext(ctx, "SELECT id, original_url, deleted FROM short_url WHERE id = $1 LIMIT 1", id)
	if err := row.Err(); err != nil {
		return nil, err
	}

	shortURL := model.ShortURL{}
	row.Scan(
		&shortURL.ID,
		&shortURL.OriginalURL,
		&shortURL.Deleted,
	)

	if shortURL.OriginalURL == "" {
		return nil, nil
	}

	return &shortURL, nil
}
