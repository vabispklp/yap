package shortener

import "context"

// Ping является прослойкой между хендлером и хранилищем
func (s *Shortener) Ping(ctx context.Context) error {
	return s.storage.Ping(ctx)
}
