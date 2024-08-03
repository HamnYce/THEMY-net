package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"server/dbhelper"
	"strconv"
)

func CreateHostHandler(w http.ResponseWriter, r *http.Request) {}

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

func UpdateHostHandler(w http.ResponseWriter, r *http.Request) {}

func DeleteHostHandler(w http.ResponseWriter, r *http.Request) {}
