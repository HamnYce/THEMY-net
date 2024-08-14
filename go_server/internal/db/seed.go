package internal_db

import (
	"database/sql"
	"os"
	"strings"
	debug "themynet/internal/debug"
)

func SeedDb(db *sql.DB) (err error) {
	debug.DebugPrintf("reading from data.sql\n")
	sqlInitTableStmt, err := os.ReadFile("data.sql")

	debug.DebugPrintf("read from data.sql\n")
	debug.DebugPrintf("head: %s\n", strings.Join(strings.Split(string(sqlInitTableStmt), "\n")[0:10], ","))

	if err != nil {
		return
	}

	_, err = db.Exec(string(sqlInitTableStmt))

	return
}
