package main

import (
	"fmt"
	"net/http"
	"os"
	handlers "themynet/api/v1/handlers"
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

func main() {
	debug.DebugPrintf("Starting server with DEBUG on")
	runConfig()
	fmt.Println("HOST: ", HOST)
	fmt.Println("PORT: ", PORT)
	fmt.Println("TURSO_DATABASE_URL: ", TURSO_DATABASE_URL)
	fmt.Println("TURSO_DATABASE TOKEN: ", TURSO_AUTH_TOKEN)

	db, err := datab.InitTursoDB(TURSO_DATABASE_URL, TURSO_AUTH_TOKEN)
	debug.CheckAndFatal(err)
	defer db.Close()

	// attach createHosts handler
	{
		debug.DebugPrintf("attaching createHost Handler\n")
		http.HandleFunc("/CreateHosts", handlers.CreateHostsHandler)
		debug.DebugPrintf("attached createHost Handler\n")
	}

	// attach RetrieveHosts handler
	{
		debug.DebugPrintf("attaching RetrieveHosts Handler\n")
		http.HandleFunc("/RetrieveHosts", handlers.RetrieveHostsHandler)
		debug.DebugPrintf("attached RetrieveHosts Handler\n")
	}

	// attach UpdateHosts handler
	{
		debug.DebugPrintf("attaching UpdateHost Handler\n")
		http.HandleFunc("/UpdateHosts", handlers.UpdateHostsHandler)
		debug.DebugPrintf("attached UpdateHost Handler\n")
	}

	// attach DeleteHosts handler
	{
		debug.DebugPrintf("attaching DeleteHost Handler\n")
		http.HandleFunc("/DeleteHosts", handlers.DeleteHostsHandler)
		debug.DebugPrintf("attached DeleteHost Handler\n")
	}

	debug.DebugPrintf("Listening on %s:%s\n", HOST, PORT)

	err = http.ListenAndServe(HOST+":"+PORT, nil)

	debug.CheckAndFatal(err)
}
