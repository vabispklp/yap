package postgres

import "context"

// Ping пингует хранилище
func (s *Storage) Ping(ctx context.Context) error {
	return s.db.PingContext(ctx)
}
