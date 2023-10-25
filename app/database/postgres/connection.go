package postgres

import (
	"database/sql"
	"fmt"
	"log/slog"

	"DataProcessorService/app/internal/config"
	_ "github.com/lib/pq"
)

func NewConnection(dbConfig config.Database) (*sql.DB, error) {
	DBConnect, err := sql.Open("postgres", dbConfig.ConnectionString())
	if err != nil {
		return nil, err
	}

	if err = DBConnect.Ping(); err != nil {
		DBConnect.Close()
		return nil, fmt.Errorf("failed to ping the database: %v", err)
	}

	slog.Info(dbConfig.ConnectionString())

	return DBConnect, nil
}
