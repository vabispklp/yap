package shortener

import "context"

func (s *Shortener) Ping(ctx context.Context) error {
	return s.db.PingContext(ctx)
}
