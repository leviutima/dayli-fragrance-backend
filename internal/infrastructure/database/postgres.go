package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func NewPostgresConnection() (*pgx.Conn, error) {
	var dsn = fmt.Sprintf(
    "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable client_encoding=UTF8",
    os.Getenv("DB_HOST"),
    os.Getenv("DB_PORT"),
    os.Getenv("DB_USER"),
    os.Getenv("DB_PASSWORD"),
    os.Getenv("DB_NAME"),
)


	var connection, err = pgx.Connect(context.Background(), dsn)
	if err != nil {
		return nil, err
	}

	return connection, nil
}
