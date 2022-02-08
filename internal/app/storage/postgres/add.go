package postgres

import (
	"context"
	"github.com/vabispklp/yap/internal/app/storage/model"
)

func (s *Storage) Add(ctx context.Context, shortURL model.ShortURL) error {
	_, err := s.db.ExecContext(ctx,
		"INSERT INTO short_url (id,user_id,original_url) VALUES ($1,$2,$3) ON CONFLICT DO NOTHING",
		shortURL.ID,
		shortURL.UserID,
		shortURL.OriginalURL,
	)

	return err
}
