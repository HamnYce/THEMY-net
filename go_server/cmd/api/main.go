package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/go-chi/cors"
	routes "themynet/api/v1/routes"
	datab "themynet/internal/db"
	debug "themynet/internal/debug"

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

	// Initialize a new ServeMux router
	mux := http.NewServeMux()

	// Pass the mux to SetupRoutes to attach routes
	routes.SetupRoutes(mux)

	// Set up CORS middleware
	cors := cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // Allow specific origin
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	// Wrap the mux with the CORS middleware
	handler := cors(mux)

	// Start the server
	debug.DebugPrintf("Listening on %s:%s\n", HOST, PORT)
	err = http.ListenAndServe(HOST+":"+PORT, handler)
	debug.CheckAndFatal(err)
}
