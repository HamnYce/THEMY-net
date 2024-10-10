package model

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"themynet/internal/db"
	debug "themynet/internal/debug"
)

const (
	TABLE_NAME = "hosts"
)

type Host struct {
	Id              *int
	Name            *string
	Ip              *string
	Mac             *string
	Hostname        *string
    Status          *int    // Didn't find anything that converts int to bool in the code soo.. 
    Exposure        *int    // Changed from *bool to *int since schema in init has them as int. 
    InternetAccess  *int    // ~khabs
	Os              *string
	OsVersion       *string
	Ports           *string
	Usage           *string
	Location        *string
	Owners          *string
	Dependencies    *string
	CreatedAt       *string
	CreatedBy       *string
	RecordedAt      *string
	Access          *string
	ConnectsTo      *string
	HostType        *string
	ExposedServices *string
	CpuCores        *int
	RamGB           *int
	StorageGB       *int
}

func (host *Host) New() {
	host.Id = new(int)
	host.Name = new(string)
	host.Ip = new(string)
	host.Mac = new(string)
	host.Hostname = new(string)
    host.Status = new(int)         // Changed to *int
    host.Exposure = new(int)       // Changed to *int
    host.InternetAccess = new(int) // Changed to *int
	host.Os = new(string)
	host.OsVersion = new(string)
	host.Ports = new(string)
	host.Usage = new(string)
	host.Location = new(string)
	host.Owners = new(string)
	host.Dependencies = new(string)
	host.CreatedAt = new(string)
	host.CreatedBy = new(string)
	host.RecordedAt = new(string)
	host.Access = new(string)
	host.ConnectsTo = new(string)
	host.HostType = new(string)
	host.ExposedServices = new(string)
	host.CpuCores = new(int)
	host.RamGB = new(int)
	host.StorageGB = new(int)
}

// Creates a host and returns the hostID
func (host Host) Createhost() (hostID int64, err error) {
	hostMap, err := host.ToMap()
	if err != nil {
		return
	}

	dbCreateStatement := "INSERT INTO " + TABLE_NAME + " "
	columns := make([]string, 0)
	values := make([]any, 0)
	for key, val := range hostMap {
		if key == "Id" {
			continue
		}
		columns = append(columns, key)
		values = append(values, val)
	}

	dbCreateStatement += fmt.Sprintf(
		"( %s ) VALUES ( %s )",
		strings.Join(columns, ","),
		strings.Join(strings.Split(strings.Repeat("?", len(columns)), ""), ","),
	)

	debug.DebugPrintf("Executing Createhost Statement\n")
	res, err := db.DBSingleton().Exec(dbCreateStatement, values...)

	if err != nil {
		return
	}

	return res.LastInsertId()
}

func Retrievehost(hostID int) (host Host, err error) {
	//dbGetStatement := "SELECT hostid, * FROM " + TABLE_NAME + " WHERE rowid = ?"
	dbGetStatement := "SELECT * FROM " + TABLE_NAME + " WHERE id = ?"  //Changed query

	debug.DebugPrintf("Executing Retrievehost Statement\n")
	sqlhost, err := db.DBSingleton().Query(dbGetStatement, hostID)
	if err != nil {
		return
	}
	defer sqlhost.Close()

	if sqlhost.Next() {
		err = host.FromsqlHosts(sqlhost)

		if err != nil {
			return
		}
	}

	return
}

func Retrievehosts(amount, offset int) (hosts []Host, err error) {
	//dbGetStatement := fmt.Sprintf("SELECT rowid, * FROM "+TABLE_NAME+" LIMIT %d OFFSET %d", amount, offset)
	dbGetStatement := fmt.Sprintf("SELECT * FROM "+TABLE_NAME+" LIMIT %d OFFSET %d", amount, offset) //Changed query

	debug.DebugPrintf("Executing Retrievehosts Statement\n")
	sqlHosts, err := db.DBSingleton().Query(dbGetStatement)
	if err != nil {
		return
	}

	debug.DebugPrintf("Executed Retrievehosts Statement\n")

	for sqlHosts.Next() {
		host := new(Host)
		err = host.FromsqlHosts(sqlHosts)
		if err != nil {
			continue
		}

		hosts = append(hosts, *host)
	}

	return
}

func (host *Host) Updatehost(hostMap map[string]any) (err error) {
	if hostMap["id"] == nil {
		err = errors.New("id is required to update a host")
		return
	}

	err = host.ParseMap(hostMap)
	if err != nil {
		return
	}

	dbUpdateStatement := "UPDATE " + TABLE_NAME + " SET "

	updates := make([]string, 0)
	values := make([]any, 0)
	for key, val := range hostMap {
		if key == "id" {
			continue
		}

		updates = append(updates, fmt.Sprintf("%s = ?", key)) // "key = ?"
		values = append(values, val)
	}
	values = append(values, *host.Id)

	dbUpdateStatement += strings.Join(updates, ", ")
	dbUpdateStatement += " WHERE rowid = ? "

	debug.DebugPrintf("Executing Updatehost Statement\n: %s\n", dbUpdateStatement)
	_, err = db.DBSingleton().Exec(dbUpdateStatement, values...)
	if err != nil {
		return
	}

	return
}

func (host Host) Deletehost() (success bool, err error) {
	dbDeleteStatement := "DELETE FROM " + TABLE_NAME + "  WHERE rowid = ?"

	debug.DebugPrintf("Executing Deletehost Statement\n")
	res, err := db.DBSingleton().Exec(dbDeleteStatement, *host.Id)
	if err != nil {
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return
	}
	return rowsAffected > 0, nil
}

// Scans current host from sql.hosts into dbhelper.host
// func (host *Host) FromsqlHosts(sqlHosts *sql.Rows) (err error) {
// 	debug.DebugPrintf("Scanning host\n")
// 	return sqlHosts.Scan(
// 		&host.Id,
// 		&host.Name,
// 		&host.Mac,
// 		&host.Ip,
// 		&host.Hostname,
// 		&host.Status,
// 		&host.Exposure,
// 		&host.InternetAccess,
// 		&host.Os,
// 		&host.OsVersion,
// 		&host.Ports,
// 		&host.Usage,
// 		&host.Location,
// 		&host.Owners,
// 		&host.Dependencies,
// 		&host.CreatedAt,
// 		&host.CreatedBy,
// 		&host.RecordedAt,
// 		&host.Access,
// 		&host.ConnectsTo,
// 		&host.HostType,
// 		&host.ExposedServices,
// 		&host.CpuCores,
// 		&host.RamGB,
// 		&host.StorageGB,
// 	)
// }


//Debug Version.
func (host *Host) FromsqlHosts(sqlHosts *sql.Rows) (err error) {
    debug.DebugPrintf("Scanning host\n")

    // Define an array of placeholders for the Scan function
    var (
        id, status, exposure, internetAccess, cpuCores, ramGB, storageGB int
        name, mac, ip, hostname, os, osVersion, ports, usage string
        location, owners, dependencies, createdAt, createdBy, recordedAt string
        access, connectsTo, hostType, exposedServices string
    )

    // Scan the fields one by one with individual error handling
    if err := sqlHosts.Scan(
        &id,
        &name,
        &mac,
        &ip,
        &hostname,
        &status,
        &exposure,
        &internetAccess,
        &os,
        &osVersion,
        &ports,
        &usage,
        &location,
        &owners,
        &dependencies,
        &createdAt,
        &createdBy,
        &recordedAt,
        &access,
        &connectsTo,
        &hostType,
        &exposedServices,
        &cpuCores,
        &ramGB,
        &storageGB,
    ); err != nil {
        debug.DebugPrintf("Error during Scan:\n")
        debug.DebugPrintf("  id: %v\n  name: %v\n  mac: %v\n  ip: %v\n  hostname: %v\n", id, name, mac, ip, hostname)
        debug.DebugPrintf("  status: %v\n  exposure: %v\n  internetAccess: %v\n", status, exposure, internetAccess)
        debug.DebugPrintf("  os: %v\n  osVersion: %v\n  ports: %v\n  usage: %v\n", os, osVersion, ports, usage)
        debug.DebugPrintf("  location: %v\n  owners: %v\n  dependencies: %v\n", location, owners, dependencies)
        debug.DebugPrintf("  createdAt: %v\n  createdBy: %v\n  recordedAt: %v\n", createdAt, createdBy, recordedAt)
        debug.DebugPrintf("  access: %v\n  connectsTo: %v\n  hostType: %v\n  exposedServices: %v\n", access, connectsTo, hostType, exposedServices)
        debug.DebugPrintf("  cpuCores: %v\n  ramGB: %v\n  storageGB: %v\n", cpuCores, ramGB, storageGB)
        debug.DebugPrintf("Scan error: %v\n", err)
        return err
    }

    // Now assign the successfully scanned values to the struct fields
    host.Id = &id
    host.Name = &name
    host.Mac = &mac
    host.Ip = &ip
    host.Hostname = &hostname
    host.Status = &status
    host.Exposure = &exposure
    host.InternetAccess = &internetAccess
    host.Os = &os
    host.OsVersion = &osVersion
    host.Ports = &ports
    host.Usage = &usage
    host.Location = &location
    host.Owners = &owners
    host.Dependencies = &dependencies
    host.CreatedAt = &createdAt
    host.CreatedBy = &createdBy
    host.RecordedAt = &recordedAt
    host.Access = &access
    host.ConnectsTo = &connectsTo
    host.HostType = &hostType
    host.ExposedServices = &exposedServices
    host.CpuCores = &cpuCores
    host.RamGB = &ramGB
    host.StorageGB = &storageGB

    return nil
}

func (host *Host) ParseMap(hostMap map[string]any) (err error) {
	jsonhost, err := json.Marshal(hostMap)
	if err != nil {
		return
	}
	err = json.Unmarshal(jsonhost, &host)
	return
}

func (host Host) ToMap() (hostMap map[string]any, err error) {
	jsonhost, err := json.Marshal(host)
	if err != nil {
		return
	}
	err = json.Unmarshal(jsonhost, &hostMap)
	return
}
