package postgres

import (
	"context"
	"github.com/vabispklp/yap/internal/app/storage/model"
)

func (s *Storage) GetByUser(ctx context.Context, userID string) ([]model.ShortURL, error) {
	var (
		item model.ShortURL
	)

	rows, err := s.db.QueryContext(ctx, "SELECT id, original_url FROM short_url where user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]model.ShortURL, 0)
	for rows.Next() {
		if err = rows.Scan(&item.ID, &item.OriginalURL); err != nil {
			return nil, err
		}

		result = append(result, item)
	}

	return result, nil
}
