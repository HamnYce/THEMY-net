package db

import (
	"database/sql"
	"errors"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

var dbRef *sql.DB

// FUTURE: add a function to close the database connection
// FUTURE: add a function to check if the database connection is open
// FUTURE: init dev or prod database based on environment variable

func InitTursoDB(url, token string) (err error) {
	if url == "" || token == "" {
		return errors.New("url and token must be provided as environment variables: TURSO_DATABASE_URL and TURSO_DATABASE_TOKEN")
	}
	db, err := sql.Open("libsql", url+"?authToken="+token)
	dbRef = db
	return
}

func DBSingleton() *sql.DB {
	return dbRef
}
