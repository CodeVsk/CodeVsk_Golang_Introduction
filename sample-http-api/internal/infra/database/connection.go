package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Connection()(*sql.DB, error) {
	conection_string := "host=localhost user=postgres password=dev@2024 dbname=Shop sslmode=disable"
	db, err := sql.Open("postgres", conection_string)

	if err != nil {
		return nil, err
	}

	//defer db.Close()

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

