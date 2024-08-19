package internal_model

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
	Status          *bool
	Exposure        *bool
	InternetAccess  *bool
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
	host.Status = new(bool)
	host.Exposure = new(bool)
	host.InternetAccess = new(bool)
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
	dbGetStatement := "SELECT hostid, * FROM " + TABLE_NAME + " WHERE rowid = ?"

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
	dbGetStatement := fmt.Sprintf("SELECT rowid, * FROM "+TABLE_NAME+" LIMIT %d OFFSET %d", amount, offset)

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
func (host *Host) FromsqlHosts(sqlHosts *sql.Rows) (err error) {
	debug.DebugPrintf("Scanning host\n")
	return sqlHosts.Scan(
		&host.Id,
		&host.Name,
		&host.Mac,
		&host.Ip,
		&host.Hostname,
		&host.Status,
		&host.Exposure,
		&host.InternetAccess,
		&host.Os,
		&host.OsVersion,
		&host.Ports,
		&host.Usage,
		&host.Location,
		&host.Owners,
		&host.Dependencies,
		&host.CreatedAt,
		&host.CreatedBy,
		&host.RecordedAt,
		&host.Access,
		&host.ConnectsTo,
		&host.HostType,
		&host.ExposedServices,
		&host.CpuCores,
		&host.RamGB,
		&host.StorageGB,
	)
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
