package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Use the appropriate driver for your database
)

func Connect(cfg DBConfig) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode)
	return sql.Open("postgres", connStr)
}
