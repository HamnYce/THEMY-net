package main

import (
	"fmt"
	"net/http"
	"os"
	routes "themynet/api/v1/routes"
	datab "themynet/internal/db"
	debug "themynet/internal/debug"

	"github.com/joho/godotenv"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

var (
	HOST               string
	PORT               string
	TURSO_DATABASE_URL string
	TURSO_AUTH_TOKEN   string
)

func runConfig() {
	godotenv.Load("./.env_local")
	HOST = os.Getenv("HOST")
	PORT = os.Getenv("PORT")
	TURSO_DATABASE_URL = os.Getenv("TURSO_DATABASE_URL")
	TURSO_AUTH_TOKEN = os.Getenv("TURSO_AUTH_TOKEN")

	if os.Getenv("DEBUG") == "true" {
		debug.SetDebug(true)
	}
}

func debugConfig() {
	fmt.Println("HOST: ", HOST)
	fmt.Println("PORT: ", PORT)
	fmt.Println("TURSO_DATABASE_URL: ", TURSO_DATABASE_URL)
	fmt.Println("TURSO_DATABASE TOKEN: ", TURSO_AUTH_TOKEN)
}

func main() {
	debug.DebugPrintf("Starting server with DEBUG on")
	runConfig()
	debugConfig()

	db, err := datab.InitTursoDB(TURSO_DATABASE_URL, TURSO_AUTH_TOKEN)
	debug.CheckAndFatal(err)
	defer db.Close()

	routes.SetupRoutes()

	debug.DebugPrintf("Listening on %s:%s\n", HOST, PORT)

	err = http.ListenAndServe(HOST+":"+PORT, nil)

	debug.CheckAndFatal(err)
}
