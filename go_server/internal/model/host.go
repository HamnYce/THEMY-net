package model

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
	utils "themynet"
	"themynet/internal/db"
	"time"
)

const (
	TABLE_NAME = "hosts"
)

// Creates a host and returns the hostID
func Createhost(host Host) (hostID int64, err error) {
	// TODO: make createdAt a DEFAULT CURRENT_DATE
	host.CreatedAt = time.Now().Unix()
	host.RecordedAt = time.Now().Unix()
	host.CreatedBy = "HamnYce"

	hostMap, err := HostToMap(host)
	if err != nil {
		return
	}

	columns := make([]string, 0)
	values := make([]any, 0)
	for key, val := range hostMap {
		if key == "Id" {
			continue
		}
		columns = append(columns, key)
		values = append(values, val)
	}

	dbCreateStatement := fmt.Sprintf(
		"INSERT INTO %s ( %s ) VALUES ( %s )",
		TABLE_NAME,
		strings.Join(columns, ","),
		strings.Join(strings.Split(strings.Repeat("?", len(columns)), ""), ","),
	)

	utils.DebugPrintf("Executing Createhost Statement\n")
	res, err := db.DBSingleton().Exec(dbCreateStatement, values...)
	if err != nil {
		return
	}

	return res.LastInsertId()
}

func Retrievehost(hostId int64) (host *Host, err error) {
	dbGetStatement := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", TABLE_NAME)

	utils.DebugPrintf("Executing Retrievehost Statement\n")

	sqlhost, err := db.DBSingleton().Query(dbGetStatement, hostId)
	if err != nil {
		return
	}

	if sqlhost.Next() {
		host, err = scanHostFromsqlRows(sqlhost)
	} else {
		err = errors.New("does not exist")
	}

	return
}

func Retrievehosts(amount, offset uint) (hosts []Host, err error) {
	dbRetrieveStatement := fmt.Sprintf("SELECT * FROM %s LIMIT ? OFFSET ?", TABLE_NAME)

	utils.DebugPrintf("Executing Retrievehosts Statement\n")
	sqlHosts, err := db.DBSingleton().Query(dbRetrieveStatement, amount, offset)
	if err != nil {
		return
	}

	utils.DebugPrintf("Executed Retrievehosts Statement\n")

	for sqlHosts.Next() {
		host, err := scanHostFromsqlRows(sqlHosts)

		if err != nil {
			break
		}

		hosts = append(hosts, *host)
	}

	return
}

func Updatehosts(hostMaps []map[string]any) (hosts []Host, err error) {
	var validIds []int64
	for _, hostMap := range hostMaps {
		var updateColumns []string
		var updateValues []any
		for column, value := range hostMap {
			if column == "Id" {
				continue
			}

			updateColumns = append(updateColumns, column+" = ?")
			updateValues = append(updateValues, value)
		}
		updateValues = append(updateValues, hostMap["Id"])

		dbUpdateStatement := fmt.Sprintf(
			"UPDATE %s SET %s WHERE id = ?",
			TABLE_NAME,
			strings.Join(updateColumns, ","),
		)

		_, err = db.DBSingleton().Exec(dbUpdateStatement, updateValues...)
		if err != nil {
			log.Println(err)
			continue
		}
		validIds = append(validIds, int64(hostMap["Id"].(float64)))
	}

	// retrieve just updated hosts
	for _, validId := range validIds {
		host, err := Retrievehost(validId)
		if err != nil {
			log.Println(err)
			continue
		}

		hosts = append(hosts, *host)
	}

	return
}

func Deletehost(hostID int64) (err error) {
	dbDeleteStatement := fmt.Sprintf("DELETE FROM %s WHERE id = ?", TABLE_NAME)

	res, err := db.DBSingleton().Exec(dbDeleteStatement, hostID)
	if err != nil {
		return
	}

	_, err = res.RowsAffected()

	return
}

// Scans current host from sql.hosts into dbhelper.host
func scanHostFromsqlRows(sqlHosts *sql.Rows) (host *Host, err error) {
	utils.DebugPrintf("Scanning host from sql row\n")
	host = &Host{}
	err = sqlHosts.Scan(
		&host.Id,
		&host.Name,
		&host.Mac,
		&host.Ip,
		&host.Ports,
		&host.Hostname,
		&host.HostType,
		&host.Os,
		&host.OsVersion,
		&host.Dependencies,
		&host.ExposedServices,
		&host.Status,
		&host.Exposure,
		&host.InternetAccess,
		&host.CpuCores,
		&host.RamGB,
		&host.StorageGB,
		&host.Usage,
		&host.Location,
		&host.Owners,
		&host.Access,
		&host.ConnectsTo,
		&host.CreatedBy,
		&host.CreatedAt,
		&host.RecordedAt,
	)

	return
}

func HostToMap(h Host) (m map[string]any, err error) {
	hostJSON, err := json.Marshal(h)
	if err != nil {
		return
	}
	err = json.Unmarshal(hostJSON, &m)
	return
}

type Host struct {
	Id              int64
	Name            *string
	Mac             *string
	Ip              *string
	Ports           *string
	Hostname        *string
	HostType        *string
	Os              *string
	OsVersion       *string
	Dependencies    *string
	ExposedServices *string
	Status          *bool
	Exposure        *bool
	InternetAccess  *bool
	CpuCores        *int
	RamGB           *int
	StorageGB       *int
	Usage           *string
	Location        *string
	Owners          *string
	Access          *string
	ConnectsTo      *string
	CreatedBy       string
	CreatedAt       int64
	RecordedAt      int64
}

func newHost() *Host {
	return &Host{
		Name:            new(string),
		Mac:             new(string),
		Ip:              new(string),
		Ports:           new(string),
		Hostname:        new(string),
		HostType:        new(string),
		Os:              new(string),
		OsVersion:       new(string),
		Dependencies:    new(string),
		ExposedServices: new(string),
		Status:          new(bool),
		Exposure:        new(bool),
		InternetAccess:  new(bool),
		CpuCores:        new(int),
		RamGB:           new(int),
		StorageGB:       new(int),
		Usage:           new(string),
		Location:        new(string),
		Owners:          new(string),
		Access:          new(string),
		ConnectsTo:      new(string),
	}
}
