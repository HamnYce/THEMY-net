package dbhelper

import (
	"database/sql"
	"os"
	"server/globalhelpers"
	"strings"
)

func SeedDb(db *sql.DB) (err error) {
	globalhelpers.DebugPrintf("reading from data.sql\n")
	sqlInitTableStmt, err := os.ReadFile("data.sql")

	globalhelpers.DebugPrintf("read from data.sql\n")
	globalhelpers.DebugPrintf("head: %s\n", strings.Join(strings.Split(string(sqlInitTableStmt), "\n")[0:10], ","))

	if err != nil {
		return
	}

	_, err = db.Exec(string(sqlInitTableStmt))

	return
}
