package postgres

import (
	"context"
)

// Delete удаляет сокращенные ссылки
func (s *Storage) Delete(ctx context.Context, ids []string, userID string) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	txStmt, err := tx.PrepareContext(ctx, "UPDATE short_url SET deleted = true WHERE id = $1 and user_id = $2")
	if err != nil {
		return err
	}

	for _, id := range ids {
		_, err = txStmt.ExecContext(ctx, id, userID)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}
