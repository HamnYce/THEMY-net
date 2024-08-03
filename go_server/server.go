package main

import (
	"database/sql"
	"log"
	"net/http"
	"server/dbhelper"

	_ "github.com/mattn/go-sqlite3"
)

const (
	DEBUG = true
	SEED  = false
)

func main() {
	if DEBUG {
		log.Println("Starting server with DEBUG on")
	}
	db, err := sql.Open("sqlite3", "data.sqlite3")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if SEED {
		log.Println("Seeding database")
		err = dbhelper.SeedDb(db)
		if err != nil {
			log.Fatal(err)
		}
	}

	if DEBUG {
		log.Println("attaching createHost Handler")
	}
	http.HandleFunc("/createHost", CreateHostHandler)

	if DEBUG {
		log.Println("attaching RetrieveHosts Handler")
	}
	http.HandleFunc("/RetrieveHosts", RetrieveHostsHandler(db))

	if DEBUG {
		log.Println("attaching UpdateHost Handler")
	}
	http.HandleFunc("/updateHost", UpdateHostHandler)

	if DEBUG {
		log.Println("attaching DeleteHost Handler")
	}
	http.HandleFunc("/deleteHost", DeleteHostHandler)

	if DEBUG {
		log.Println("Listening on port 8091")
	}
	err = http.ListenAndServe(":8091", nil)

	if err != nil {
		log.Fatal(err)
	}
}

// NOTE: deleting rows
// err = dbhelper.DeleteRow(db, 1)
// if err != nil {
// 	log.Fatal(err)
// }

// NOTE: retrieving rows
// rows, err := dbhelper.RetrieveRows(db, 1, 52)

// if err != nil {
// 	log.Fatal(err)
// }

// for _, row := range rows {
// 	pretty.Println(row)
// }

// NOTE: updating rows
// rowMap := make(map[string]any)
// rowMap["Id"] = 2
// rowMap["Name"] = "untest"
// rowMap["Ip"] = "69.69"
// rowsAffected, err := dbhelper.UpdateRow(db, rowMap)
// if err != nil {
// 	log.Fatal(err)
// }
// log.Println("rows affected: ", rowsAffected)

/*
	// NOTE: creating rows
	row := new(dbhelper.Row)
	name := "test"
	ip := "1.1.1.1"
	hostname := "hostname"
	// status := true
	exposure := true
	internetAccess := false
	// os := "lunix"
	osVersion := "1.2.2"
	ports := "[1,2,3,4]"
	usage := "[dataserver, webserver]"
	location := "location"
	owners := "[owner1, owner2]"
	// dependencies := "[dep1, dep2]"
	createdAt := "2021-01-01"
	createdBy := "admin"
	recordedAt := "2021-01-01"
	// access := "public"
	connectsTo := "connectsTo"
	hostType := "hostType"
	exposedServices := "[service1, service2]"
	cpuCores := 20
	ramGB := 32
	storageGB := 1024

	row.Name = &name
	row.Ip = &ip
	row.Hostname = &hostname
	row.Status = nil
	row.Exposure = &exposure
	row.InternetAccess = &internetAccess
	row.Os = nil
	row.OsVersion = &osVersion
	row.Ports = &ports
	row.Usage = &usage
	row.Location = &location
	row.Owners = &owners
	row.Dependencies = nil
	row.CreatedAt = &createdAt
	row.CreatedBy = &createdBy
	row.RecordedAt = &recordedAt
	row.Access = nil
	row.ConnectsTo = &connectsTo
	row.HostType = &hostType
	row.ExposedServices = &exposedServices
	row.CpuCores = &cpuCores
	row.RamGB = &ramGB
	row.StorageGB = &storageGB

	id, err := dbhelper.CreateRow(db, *row)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(id)
*/
