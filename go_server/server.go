package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"server/dbhelper"
	"server/globalhelpers"

	"github.com/joho/godotenv"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

const (
	HOST_IP = "100.68.33.2"
	PORT    = "8091"
	SEED    = false
)

func getDatabaseURL() (url string, err error) {
	godotenv.Load(".env.local")
	databaseUrl := os.Getenv("TURSO_DATABASE_URL")
	databaseToken := os.Getenv("TURSO_AUTH_TOKEN")

	globalhelpers.CheckAndFatal(errors.New("TURSO_DATABASE_URL not set"))
	globalhelpers.CheckAndFatal(errors.New("TURSO_AUTH_TOKEN not set"))

	url = fmt.Sprintf("%s?authToken=%s",
		databaseUrl,
		databaseToken,
	)

	return
}

func main() {

	globalhelpers.DebugPrintf("Starting server with DEBUG on")

	url, err := getDatabaseURL()
	globalhelpers.CheckAndFatal(err)

	db, err := sql.Open("libsql", url)
	globalhelpers.CheckAndFatal(err)
	defer db.Close()

	if SEED {
		log.Println("Seeding database")
		err = dbhelper.SeedDb(db)
		globalhelpers.CheckAndFatal(err)
	}

	globalhelpers.DebugPrintf("attaching createHost Handler\n")
	http.HandleFunc("/CreateHosts", CreateHostsHandler(db))
	globalhelpers.DebugPrintf("attached createHost Handler\n")

	globalhelpers.DebugPrintf("attaching RetrieveHosts Handler\n")
	http.HandleFunc("/RetrieveHosts", RetrieveHostsHandler(db))
	globalhelpers.DebugPrintf("attached RetrieveHosts Handler\n")

	globalhelpers.DebugPrintf("attaching UpdateHost Handler\n")
	http.HandleFunc("/UpdateHosts", UpdateHostsHandler(db))
	globalhelpers.DebugPrintf("attached UpdateHost Handler\n")

	globalhelpers.DebugPrintf("attaching DeleteHost Handler\n")
	http.HandleFunc("/DeleteHosts", DeleteHostsHandler(db))
	globalhelpers.DebugPrintf("attached DeleteHost Handler\n")

	globalhelpers.DebugPrintf("Listening on %s:%s\n", HOST_IP, PORT)

	err = http.ListenAndServe(HOST_IP+":"+PORT, nil)

	globalhelpers.CheckAndFatal(err)
}
