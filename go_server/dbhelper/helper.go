package dbhelper

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"server/globalhelpers"
	"strings"
)

const (
	DEBUG = false
)

// TODO: make CRUD functions into receiver methods for Row struct

type Row struct {
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

func FetchRowCount(db *sql.DB) (rowCount int) {
	globalhelpers.DebugPrintf("Executing FetchRowCount DB Statement")
	row := db.QueryRow("SELECT COUNT(*) FROM data")

	row.Scan(&rowCount)

	return
}

// Creates a row and returns the rowID
func CreateRow(db *sql.DB, row Row) (rowID int64, err error) {
	rowMap, err := RowToMap(row)
	if err != nil {
		return
	}

	dbCreateStatement := "INSERT INTO data "
	columns := make([]string, 0)
	values := make([]any, 0)
	for key, val := range rowMap {
		if key == "Id" {
			continue
		}
		columns = append(columns, key)
		values = append(values, val)
	}

	dbCreateStatement += fmt.Sprintf(
		"(%s) VALUES (%s)",
		strings.Join(columns, ","),
		strings.Join(strings.Split(strings.Repeat("?", len(columns)), ""), ","),
	)

	if DEBUG {
		log.Println("CreateRow: ", dbCreateStatement)
	}

	globalhelpers.DebugPrintf("Executing CreateRow Statement\n")
	res, err := db.Exec(dbCreateStatement, values...)

	if err != nil {
		return
	}

	return res.LastInsertId()
}

func RetrieveRow(db *sql.DB, rowID int) (row Row, err error) {
	dbGetStatement := "SELECT rowid, * FROM data WHERE rowid = ?"

	globalhelpers.DebugPrintf("Executing RetrieveRow Statement\n")
	sqlRow, err := db.Query(dbGetStatement, rowID)
	sqlRow.Next()
	defer sqlRow.Close()

	if err != nil {
		return
	}

	err = ScanRow(sqlRow, &row)

	return
}

// TODO: abstract out the sql statement generation to make filtering the rows easier
// or add map[string]any filter argument and just keep it empty if not needed
// retrieve amount rows from db starting from given offset. if amount == 1, offset acts like index
func RetrieveRows(db *sql.DB, amount, offset int) (rows []Row, err error) {
	dbGetStatement := fmt.Sprintf("SELECT rowid, * FROM data LIMIT %d OFFSET %d", amount, offset)

	globalhelpers.DebugPrintf("Executing RetrieveRows Statement\n")
	sqlRows, err := db.Query(dbGetStatement)

	for sqlRows.Next() {
		row := new(Row)
		err = ScanRow(sqlRows, row)
		if err != nil {
			return
		}

		rows = append(rows, *row)
	}

	return
}

func UpdateRow(db *sql.DB, rowMap map[string]any) (err error) {
	if rowMap["Id"] == nil {
		err = errors.New("id is required to update a row")
		return
	}

	dbUpdateStatement := "UPDATE data SET "
	updates := make([]string, 0)
	values := make([]any, 0)

	for key, val := range rowMap {
		if key == "Id" {
			continue
		}

		updates = append(updates, fmt.Sprintf("%s = ?", key))
		values = append(values, val)
	}
	values = append(values, rowMap["Id"])

	dbUpdateStatement += strings.Join(updates, ", ")
	dbUpdateStatement += " WHERE rowid = ?"

	globalhelpers.DebugPrintf("Executing UpdateRow Statement\n")
	_, err = db.Exec(dbUpdateStatement, values...)
	if err != nil {
		return
	}
	return
}

func DeleteRow(db *sql.DB, rowID int) (success bool, err error) {
	dbDeleteStatement := "DELETE FROM data WHERE rowid = ?"

	globalhelpers.DebugPrintf("Executing DeleteRow Statement\n")
	res, err := db.Exec(dbDeleteStatement, rowID)
	if err != nil {
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return
	}
	return rowsAffected > 0, nil
}

// Scans current row from sql.Rows into dbhelper.Row
func ScanRow(sqlRows *sql.Rows, row *Row) (err error) {
	err = sqlRows.Scan(
		&row.Id,
		&row.Name,
		&row.Mac,
		&row.Ip,
		&row.Hostname,
		&row.Status,
		&row.Exposure,
		&row.InternetAccess,
		&row.Os,
		&row.OsVersion,
		&row.Ports,
		&row.Usage,
		&row.Location,
		&row.Owners,
		&row.Dependencies,
		&row.CreatedAt,
		&row.CreatedBy,
		&row.RecordedAt,
		&row.Access,
		&row.ConnectsTo,
		&row.HostType,
		&row.ExposedServices,
		&row.CpuCores,
		&row.RamGB,
		&row.StorageGB,
	)

	if err != nil {
		return
	}

	return
}

func MapToRow(rowMap map[string]any) (row Row, err error) {
	jsonRow, err := json.Marshal(rowMap)
	if err != nil {
		return
	}
	err = json.Unmarshal(jsonRow, &row)
	return
}

func RowToMap(row Row) (rowMap map[string]any, err error) {
	jsonRow, err := json.Marshal(row)
	if err != nil {
		return
	}
	err = json.Unmarshal(jsonRow, &rowMap)
	return
}
