package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"server/dbhelper"
	"strconv"
)

// TODO: create host handler shoudl be pluralized
func CreateHostsHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Del("Content-Type")
		w.Header().Add("Content-Type", "application/json")
		reqJSON := make([]byte, r.ContentLength)
		resMap := make(map[string]any, 0)
		resJSON := make([]byte, 0)
		errors := make([]string, 0)
		var err error

		if r.Method != http.MethodPost {
			errors = append(errors, "Only POST method is allowed")
		}

		// reading request body
		n, err := r.Body.Read(reqJSON)
		if n == 0 && err != nil {
			log.Fatal(err)
		}

		// parsing request body into resMap
		err = json.Unmarshal(reqJSON, &resMap)
		if err != nil {
			errors = append(errors, "Invalid JSON")
		}

		if len(errors) == 0 {
			for _, v := range resMap["rows"].([]any) {
				row, err := dbhelper.MapToRow(v.(map[string]any))
				if err != nil {
					continue
				}
				dbhelper.CreateRow(db, row)
			}
		}

		resMap["errors"] = errors
		// converting resMap to resJSON
		resJSON, err = json.Marshal(resMap)
		if err != nil {
			log.Fatal(err)
		}
		w.Write(resJSON)
	}
}

func RetrieveHostsHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Del("Content-Type")
		w.Header().Add("Content-Type", "application/json")
		resMap := make(map[string]any, 0)
		errors := make([]string, 0)
		limit := 0
		offset := 0
		var err error

		if r.URL.Query().Has("limit") && r.URL.Query().Has("offset") {
			limit, err = strconv.Atoi(r.URL.Query().Get("limit"))
			if err != nil {
				errors = append(errors, "Invalid limit")
			} else {
				if limit < 0 {
					errors = append(errors, "limit has to be greater or equal to 0")
				} else {
					resMap["limit"] = limit
				}
			}

			offset, err = strconv.Atoi(r.URL.Query().Get("offset"))
			if err != nil {
				errors = append(errors, "Invalid offset")
			} else {
				if offset < 0 {
					errors = append(errors, "offset has to be greater or equal to 0")
				} else {
					resMap["offset"] = offset
				}
			}
		} else {
			errors = append(errors, "Missing limit or offset")
		}

		if r.Method != http.MethodGet {
			errors = append(errors, "Invalid method")
		}

		if len(errors) == 0 {
			rows, err := dbhelper.RetrieveRows(db, limit, offset)
			if err != nil {
				errors = append(errors, "Error retrieving rows")
			} else {
				resMap["rows"] = rows
			}
		}

		resMap["errors"] = errors
		resJSON, err := json.Marshal(resMap)
		if err != nil {
			log.Fatal(err)
		}

		w.Write(resJSON)
	}
}

// TODO: updateHost handler should be pluralized and refactored to work with multiple rows
func UpdateHostsHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Del("Content-Type")
		w.Header().Add("Content-Type", "application/json")
		reqJSON := make([]byte, r.ContentLength)
		resMap := make(map[string]any, 0)
		resJSON := make([]byte, 0)
		errors := make([]string, 0)
		var err error

		if r.Method != http.MethodPost {
			errors = append(errors, "Only POST method is allowed")
		}

		// reading request body
		n, err := r.Body.Read(reqJSON)
		if n == 0 && err != nil {
			log.Fatal(err)
		}

		// parsing request body into resMap
		err = json.Unmarshal(reqJSON, &resMap)
		if err != nil {
			errors = append(errors, "Invalid JSON")
		}

		if len(errors) == 0 {
			updatedRows := make([]map[string]any, 0)
			for _, row := range resMap["rows"].([]any) {
				err = dbhelper.UpdateRow(db, row.(map[string]any))
				if err != nil {
					errors = append(errors, err.Error())
					continue
				}
				rowID := int(row.(map[string]any)["Id"].(float64))
				row, err := dbhelper.RetrieveRow(db, rowID)
				if err != nil {
					continue
				}

				updatedRow, err := dbhelper.RowToMap(row)
				if err != nil {
					continue
				}

				updatedRows = append(updatedRows, updatedRow)
			}
			resMap["rows"] = updatedRows
		}

		resMap["errors"] = errors
		// converting resMap to resJSON
		resJSON, err = json.Marshal(resMap)
		if err != nil {
			log.Fatal(err)
		}
		w.Write(resJSON)
	}
}

func DeleteHostsHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Del("Content-Type")
		w.Header().Add("Content-Type", "application/json")
		reqJSON := make([]byte, r.ContentLength)
		resMap := make(map[string]any, 0)
		resJSON := make([]byte, 0)
		errors := make([]string, 0)
		var err error

		if r.Method != http.MethodPost {
			errors = append(errors, "Only Post method is allowed")
		}

		// reading request body
		n, err := r.Body.Read(reqJSON)
		if n == 0 && err != nil {
			log.Fatal(err)
		}

		// parsing request body into resMap
		err = json.Unmarshal(reqJSON, &resMap)
		if err != nil {
			errors = append(errors, "Invalid JSON")
		}

		deletedRowIDs := make([]int, 0)
		if len(errors) == 0 {
			for _, rawRowID := range resMap["rowIDs"].([]any) {
				rowID := int(rawRowID.(float64))
				success, err := dbhelper.DeleteRow(db, rowID)
				if err != nil {
					errors = append(errors, err.Error())
					continue
				}

				if success {
					deletedRowIDs = append(deletedRowIDs, rowID)
				}
			}
		}
		resMap["errors"] = errors
		resMap["deletedRowIDs"] = deletedRowIDs
		// converting resMap to resJSON
		resJSON, err = json.Marshal(resMap)
		if err != nil {
			log.Fatal(err)
		}
		w.Write(resJSON)
	}

}
