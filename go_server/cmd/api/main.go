package main

import (
	"net/http"
	utils "themynet"
	routes "themynet/api/v1/routes"
	"themynet/internal/db"
	datab "themynet/internal/db"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func main() {
	utils.DebugPrintf("Starting server with DEBUG on")
	utils.RunConfig()
	utils.DebugConfig()

	err := datab.InitTursoDB(utils.GetTursoURL(), utils.GetTursoAuth())
	utils.CheckAndFatal(err)
	defer db.DBSingleton().Close()

	routes.SetupRoutes()

	utils.DebugPrintf("Listening on %s:%s\n", utils.GetHost(), utils.GetPort())

	err = http.ListenAndServe(utils.GetHost()+":"+utils.GetPort(), nil)

	utils.CheckAndFatal(err)
}
