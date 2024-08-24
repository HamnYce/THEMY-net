package main

import (
	"net/http"
	utils "themynet"
	routes "themynet/api/v1/routes"
	"themynet/internal/db"
	datab "themynet/internal/db"

	"github.com/go-chi/chi/v5"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func main() {

	utils.DebugPrintf("Starting server with DEBUG on")
	utils.RunConfig()
	utils.DebugConfig()

	err := datab.InitTursoDB(utils.GetTursoURL(), utils.GetTursoAuth())
	utils.CheckAndFatal(err)
	defer db.DBSingleton().Close()

	r := chi.NewRouter()
	hostRouter := routes.SetupRoutes()

	r.Mount("/api/v1", hostRouter)

	utils.DebugPrintf("Listening on %s:%s\n", utils.GetHost(), utils.GetPort())
	err = http.ListenAndServe(utils.GetHost()+":"+utils.GetPort(), r)

	utils.CheckAndFatal(err)
}
