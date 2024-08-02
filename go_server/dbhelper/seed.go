package dbhelper

import (
	"database/sql"
	"os"
)

func SeedDb(db *sql.DB) (err error) {
	sqlInitTable, err := os.ReadFile("data.sql")
	if err != nil {
		return
	}

	_, err = db.Exec(string(sqlInitTable))

	return
}
