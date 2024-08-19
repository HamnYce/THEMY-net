package db

import (
	"database/sql"
	"errors"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

var dbRef *sql.DB

func InitTursoDB(url, token string) (db *sql.DB, err error) {
	if url == "" || token == "" {
		return nil, errors.New("url and token must be provided as environment variables: TURSO_DATABASE_URL and TURSO_DATABASE_TOKEN")
	}
	db, err = sql.Open("libsql", url+"?authToken="+token)
	dbRef = db
	return
}

func DBSingleton() *sql.DB {
	return dbRef
}
