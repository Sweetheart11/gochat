package storage

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type DB struct {
	db *sql.DB
}

func NewDB() (*DB, error) {
	conStr, err := PostgresConnStr()
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("postgres", conStr)
	if err != nil {
		return nil, err
	}

	return &DB{db: db}, nil
}

func PostgresConnStr() (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", fmt.Errorf("error loading .env file: %v\n", err)
	}

	host := os.Getenv("POSTGRESDB_HOST")
	port := os.Getenv("POSTGRESDB_PORT")
	password := os.Getenv("POSTGRESDB_PASSWORD")
	dbname := os.Getenv("POSTGRESDB_NAME")
	user := os.Getenv("POSTGRESDB_USER")
	sslmode := os.Getenv("POSTGRESDB_SSLMODE")

	connStr := fmt.Sprintf(
		"host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		host,
		port,
		user,
		password,
		dbname,
		sslmode,
	)

	return connStr, nil
}

func (d *DB) Close() {
	d.db.Close()
}

func (d *DB) GetDB() *sql.DB {
	return d.db
}
