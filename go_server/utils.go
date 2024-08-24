package public

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	debug              = false
	seed               = false
	host               string
	port               string
	turso_database_url string
	turso_auth_token   string
)

func DebugPrintf(format string, args ...any) {
	if debug {
		log.Printf(format, args...)
	}
}

func CheckAndFatal(err error) {
	if err == nil {
		return
}	

	log.Fatal(err)
}

func RunConfig() {
	godotenv.Load("./.env_local")
	host = os.Getenv("HOST")
	port = os.Getenv("PORT")
	turso_database_url = os.Getenv("TURSO_DATABASE_URL")
	turso_auth_token = os.Getenv("TURSO_AUTH_TOKEN")

	if os.Getenv("DEBUG") == "true" {
		debug = true
	}
}

func DebugConfig() {
	fmt.Println("HOST: ", host)
	fmt.Println("PORT: ", port)
	fmt.Println("TURSO_DATABASE_URL: ", turso_database_url)
	fmt.Println("TURSO_DATABASE TOKEN: ", turso_auth_token)
}

func GetTursoURL() string {
	return turso_database_url
}

func GetTursoAuth() string {
	return turso_auth_token
}

func GetHost() string {
	return host
}

func GetPort() string {
	return port
}
