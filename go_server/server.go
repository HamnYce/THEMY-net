package main

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"server/dbhelper"
	"server/globalhelpers"

	"github.com/joho/godotenv"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

var (
	HOST_IP string
	PORT    string
)

func getDatabaseURL() (url string, err error) {
	godotenv.Load(".env.local")
	databaseUrl := os.Getenv("TURSO_DATABASE_URL")
	databaseToken := os.Getenv("TURSO_AUTH_TOKEN")

	if databaseUrl == "" {
	globalhelpers.CheckAndFatal(errors.New("TURSO_DATABASE_URL not set"))
	}

	if databaseToken == "" {
	globalhelpers.CheckAndFatal(errors.New("TURSO_AUTH_TOKEN not set"))
	}

	url = fmt.Sprintf("%s?authToken=%s",
		databaseUrl,
		databaseToken,
	)

	return
}

func processFlags() {
	flag.BoolVar(&globalhelpers.DEBUG, "D", false, "Enable debug mode")
	flag.BoolVar(&globalhelpers.SEED, "seed", false, "Enable permanent seeding of the database")
	flag.StringVar(&HOST_IP, "host", "127.0.0.1", "Host IP to listen on")
	flag.StringVar(&PORT, "port", "8080", "Port to listen on")
	flag.Parse()
}

func main() {
	// TODO: test with turso
	processFlags()

	globalhelpers.DebugPrintf("Starting server with DEBUG on")

	url, err := getDatabaseURL()
	globalhelpers.CheckAndFatal(err)

	db, err := sql.Open("libsql", url)
	globalhelpers.CheckAndFatal(err)
	defer db.Close()

	if globalhelpers.SEED {
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
