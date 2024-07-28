package main

import (
	"log"
	"net/http"
	"server/gomap"
	"strings"
)

func main() {
	http.HandleFunc("/infoon", infoOnHandler)

	log.Fatal(http.ListenAndServe(":8020", nil))

}

func infoOnHandler(w http.ResponseWriter, r *http.Request) {
	log.Print(r.URL.Query())
	log.Print(r.Header)

	w.Header().Del("Content-Type")
	w.Header().Add("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		ips := strings.Split(r.URL.Query().Get("ips"), ",")

		scanner, err := gomap.NewScanner(ips)

		if err != nil {
			log.Println(err)
		}

		_, jsn, warnings, err := gomap.ScanHosts(scanner)

		if len(warnings) > 0 {
			log.Println(warnings)
		}

		if err != nil {
			log.Println(err)
		}

		_, err = w.Write(jsn)

		if err != nil {
			log.Print(err)
		}

	default:
		log.Print("Method was not get request")
		w.Write([]byte("{error: \"only accepting get requests\"}"))

	}
}
