package postgres

import (
	"context"
	"database/sql"
	"time"
)

const initSQL = "CREATE TABLE IF NOT EXISTS short_url (" +
	"id VARCHAR(100) NOT NULL," +
	"user_id VARCHAR(16) NOT NULL," +
	"original_url VARCHAR(100) NOT NULL," +
	"PRIMARY KEY (id)" +
	");"

type Storage struct {
	db *sql.DB
}

func NewStorage(dsn string) (*Storage, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	pingCTX, pingCancel := context.WithTimeout(ctx, 1*time.Second)
	defer pingCancel()
	if err = db.PingContext(pingCTX); err != nil {
		return nil, err
	}

	initCTX, initCancel := context.WithTimeout(ctx, 20*time.Second)
	defer initCancel()
	_, err = db.ExecContext(initCTX, initSQL)
	if err != nil {
		return nil, err
	}

	return &Storage{
		db: db,
	}, nil
}

func (s *Storage) Close() error {
	return s.db.Close()
}
