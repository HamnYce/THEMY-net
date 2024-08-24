package main

import (
	"bytes"
	"log"
	"os"
	utils "themynet"
	"themynet/internal/db"
)

// TODO: this requires you to manually go back in migration i.e. atm delete the entire hosts table
func main() {
	utils.RunConfig()
	utils.DebugConfig()
	db.InitTursoDB(utils.GetTursoURL(), utils.GetTursoAuth())

	sqlCreateTableStmt, err := os.ReadFile("cmd/db/migrations/2024-08-15_init.sql")
	if err != nil {
		log.Fatal(err)
	}

	sqlSeedTableStmt, err := os.ReadFile("cmd/db/seeds/2024-08-15_init_seed.sql")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.DBSingleton().Exec(string(
		bytes.Join(
			[][]byte{sqlCreateTableStmt, sqlSeedTableStmt}, []byte(";\n"),
		),
	))
	if err != nil {
		log.Fatal(err)
	}
}
