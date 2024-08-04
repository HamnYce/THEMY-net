package main

import (
	"database/sql"
	"log"
	"net/http"
	"server/dbhelper"

	_ "github.com/mattn/go-sqlite3"
)

const (
	DEBUG = true
	SEED  = false
)

func main() {
	if DEBUG {
		log.Println("Starting server with DEBUG on")
	}
	db, err := sql.Open("sqlite3", "data.sqlite3")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if SEED {
		log.Println("Seeding database")
		err = dbhelper.SeedDb(db)
		if err != nil {
			log.Fatal(err)
		}
	}

	if DEBUG {
		log.Println("attaching createHost Handler")
	}
	http.HandleFunc("/CreateHosts", CreateHostsHandler(db))

	if DEBUG {
		log.Println("attaching RetrieveHosts Handler")
	}
	http.HandleFunc("/RetrieveHosts", RetrieveHostsHandler(db))

	if DEBUG {
		log.Println("attaching UpdateHost Handler")
	}
	http.HandleFunc("/UpdateHosts", UpdateHostsHandler(db))

	if DEBUG {
		log.Println("attaching DeleteHost Handler")
	}
	http.HandleFunc("/DeleteHosts", DeleteHostsHandler(db))

	if DEBUG {
		log.Println("Listening on port 8091")
	}
	err = http.ListenAndServe(":8091", nil)

	if err != nil {
		log.Fatal(err)
	}
}
