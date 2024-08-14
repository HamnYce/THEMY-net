package cmd_api

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	handlers "themynet/api/v1/handlers"
  debug "themynet/internal/debug"
  dbhelper "themynet/internal/db"

	"github.com/BurntSushi/toml"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

var (
	HOST_IP string
	PORT    string
)

func getTursoURLFromToml() (url string, err error) {
	var tomlMap map[string]any
	toml.DecodeFile("env.toml", &tomlMap)

	if tomlMap["TURSO"] == nil {
		return url, errors.New("Turso not in env.toml")
	}
	tursoMap := tomlMap["TURSO"].(map[string]any)

	if tursoMap["TURSO_DATABASE_URL"] == nil {
		debug.CheckAndFatal(errors.New("TURSO_DATABASE_URL not set"))
	}

	if tursoMap["TURSO_AUTH_TOKEN"] == nil {
		debug.CheckAndFatal(errors.New("TURSO_AUTH_TOKEN not set"))
	}

	url = fmt.Sprintf("%s?authToken=%s",
		tursoMap["TURSO_DATABASE_URL"],
		tursoMap["TURSO_AUTH_TOKEN"],
	)

	return
}

func configWithToml() {
  // TODO: get information about debug and such from the toml file
  //  set debug and seed here
  log.Fatal("Implement this. at the moment its just setting debug debug for the api section")
}

func Main() {
	// FIXME: test with turso
	debug.DebugPrintf("Starting server with DEBUG on")

	url, err := getTursoURLFromToml()
	debug.CheckAndFatal(err)

	db, err := sql.Open("libsql", url)
	debug.CheckAndFatal(err)
	defer db.Close()

	if debug.SEED {
		log.Println("Seeding database")
		err = dbhelper.SeedDb(db)
		debug.CheckAndFatal(err)
	}

	// attach createHosts handler
	{
		debug.DebugPrintf("attaching createHost Handler\n")
		http.HandleFunc("/CreateHosts", handlers.CreateHostsHandler(db))
		debug.DebugPrintf("attached createHost Handler\n")
	}

	// attach RetrieveHosts handler
	{
		debug.DebugPrintf("attaching RetrieveHosts Handler\n")
		http.HandleFunc("/RetrieveHosts", handlers.RetrieveHostsHandler(db))
		debug.DebugPrintf("attached RetrieveHosts Handler\n")
	}

	// attach UpdateHosts handler
	{
		debug.DebugPrintf("attaching UpdateHost Handler\n")
		http.HandleFunc("/UpdateHosts", handlers.UpdateHostsHandler(db))
		debug.DebugPrintf("attached UpdateHost Handler\n")
	}

	// attach DeleteHosts handler
	{
		debug.DebugPrintf("attaching DeleteHost Handler\n")
		http.HandleFunc("/DeleteHosts", handlers.DeleteHostsHandler(db))
		debug.DebugPrintf("attached DeleteHost Handler\n")
	}

	debug.DebugPrintf("Listening on %s:%s\n", HOST_IP, PORT)

	err = http.ListenAndServe(HOST_IP+":"+PORT, nil)

	debug.CheckAndFatal(err)
}
