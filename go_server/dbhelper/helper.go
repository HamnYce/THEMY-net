package dbhelper

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
)

const (
	DEBUG = false
)

type Row struct {
	Id              *int
	Name            *string
	Ip              *string
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

	res, err := db.Exec(dbCreateStatement, values...)

	if err != nil {
		return
	}

	return res.LastInsertId()
}

// TODO: abstract out the sql statement generation to make filtering the rows easier
// or add map[string]any filter argument and just keep it empty if not needed
// retrieve amount rows from db starting from given offset. if amount == 1, offset acts like index
func RetrieveRows(db *sql.DB, amount, offset int) (rows []Row, err error) {
	dbGetStatement := fmt.Sprintf("SELECT rowid, * FROM data LIMIT %d OFFSET %d", amount, offset)

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

func UpdateRow(db *sql.DB, rowMap map[string]any) (rowsAffected int64, err error) {
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

	res, err := db.Exec(dbUpdateStatement, values...)
	if err != nil {
		return
	}
	return res.RowsAffected()
}

func DeleteRow(db *sql.DB, rowID int) (err error) {
	dbDeleteStatement := "DELETE FROM data WHERE rowid = ?"
	_, err = db.Exec(dbDeleteStatement, rowID)
	return
}

// Scans current row from sql.Rows into dbhelper.Row
func ScanRow(sqlRows *sql.Rows, row *Row) (err error) {
	err = sqlRows.Scan(
		&row.Id,
		&row.Name,
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
