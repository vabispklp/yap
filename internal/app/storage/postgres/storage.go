package postgres

import (
	"context"
	"database/sql"
	"github.com/vabispklp/yap/internal/app/storage"
	"time"
)

const initSQL = "CREATE TABLE IF NOT EXISTS short_url (" +
	"id VARCHAR(100) NOT NULL," +
	"user_id VARCHAR(32) NOT NULL," +
	"original_url VARCHAR(100) NOT NULL," +
	"deleted BOOLEAN NOT NULL DEFAULT FALSE," +
	"PRIMARY KEY (id,user_id)" +
	");"

type Storage struct {
	db *sql.DB
}

func NewStorage(dsn string) (storage.StorageExpected, error) {
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

// Close закрывает соединение с хранилищем
func (s *Storage) Close() error {
	return s.db.Close()
}
