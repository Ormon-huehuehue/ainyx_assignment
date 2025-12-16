package repository

import (
	"database/sql"
	"fmt"

	"os"

	"go-backend-task/db/sqlc"

	_ "github.com/lib/pq"
)

type Repository struct {
	Queries *db.Queries
	DB      *sql.DB
}

func NewRepository() (*Repository, error) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPass, dbName)

	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	if err := conn.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &Repository{
		Queries: db.New(conn),
		DB:      conn,
	}, nil
}

func (r *Repository) Close() {
	if r.DB != nil {
		r.DB.Close()
	}
}
