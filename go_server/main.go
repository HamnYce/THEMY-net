package main

import (
	"flag"
	"log"
	api "themynet/cmd/api"
)

func main() {
	var cmd string
  flag.StringVar(&cmd, "cmd", "", "command to run options: api/tui")
	flag.Parse()
	switch cmd {
	case "api":
		api.Main()
  case "tui":
  log.Fatal("Still not there yet ;)!")
	default:
		log.Fatal("Invalid command. use flag -h to get help")
	}
	flag.Parse()
}
