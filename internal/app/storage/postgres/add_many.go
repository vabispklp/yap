package postgres

import (
	"context"
	"github.com/vabispklp/yap/internal/app/storage/model"
)

func (s *Storage) AddMany(ctx context.Context, shortURLs []model.ShortURL) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	txStmt, err := tx.PrepareContext(ctx, "INSERT INTO short_url (id,user_id,original_url) VALUES ($1,$2,$3) ON CONFLICT DO NOTHING")
	if err != nil {
		return err
	}

	for _, shortURL := range shortURLs {
		_, err = txStmt.ExecContext(ctx,
			shortURL.ID,
			shortURL.UserID,
			shortURL.OriginalURL,
		)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}
