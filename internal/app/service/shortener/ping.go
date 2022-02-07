package shortener

import "context"

func (s *Shortener) Ping(ctx context.Context) error {
	return s.storage.Ping(ctx)
}
