package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func GetPostgresDB() (*sql.DB, error) {
	host := "localhost"
	user := "postgres"
	password := "root"
	databaseName := "golangcrud"

	desc := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, databaseName)

	db, err := createConnection(desc)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func createConnection(desc string) (*sql.DB, error) {
	db, err := sql.Open("postgres", desc)
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	return db, nil
}
