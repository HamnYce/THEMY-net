package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"server/dbhelper"
)

// struct to hold variables used in all CRUD handlers
type jSONReqResVars struct {
	reqJSON []byte
	resMap  map[string]any
	resJSON []byte
	errors  []string
	err     error
}

func newJSONReqResVars() *jSONReqResVars {
	j := new(jSONReqResVars)
	j.reqJSON = make([]byte, 0)
	j.resMap = make(map[string]any, 0)
	j.resJSON = make([]byte, 0)
	j.errors = make([]string, 0)
	j.err = nil

	return j
}

// readReqBody reads request body into jSONReqResVars.reqJSON
// stores any errors in jSONReqResVars.errors
func (j *jSONReqResVars) readReqBody(r *http.Request) {
	var n int

	j.reqJSON = make([]byte, r.ContentLength)
	n, j.err = r.Body.Read(j.reqJSON)
	if n == 0 && j.err != nil {
		j.errors = append(j.errors, j.err.Error())
	} else {
		j.err = nil
	}
}

// parseResMapIntoResJSON converts reqJSON into resMap
// stores any errors in jSONReqResVars.errors
func (j *jSONReqResVars) parseReqJSON() {
	j.err = json.Unmarshal(j.reqJSON, &j.resMap)
	if j.err != nil {
		j.errors = append(j.errors, "Invalid JSON")
	}
}

func (j *jSONReqResVars) readAndParseReqJSON(r *http.Request) {
	j.readReqBody(r)
	if j.err != nil {
		j.reqJSON = []byte("{}")
	}
	j.parseReqJSON()
}

func (j *jSONReqResVars) jSONifyResMap() {
	j.resMap["errors"] = j.errors
	j.resJSON, j.err = json.Marshal(j.resMap)
	if j.err != nil {
		j.resJSON = []byte("{}")
	}
}

func setHeaderContentTypeJSON(w *http.ResponseWriter) {
	(*w).Header().Del("Content-Type")
	(*w).Header().Add("Content-Type", "application/json")
}

// takes in json with ["rows"] key containing array of rows to create
// returns json with ["rows"] key containing array of created rows with all attributes
// if any error occurs, ["errors"] key will contain array of error messages
func CreateHostsHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		setHeaderContentTypeJSON(&w)
		j := newJSONReqResVars()

		if r.Method != http.MethodPost {
			j.errors = append(j.errors, "Only POST method is allowed")
		}
		j.readAndParseReqJSON(r)

		if len(j.errors) == 0 {
			for _, v := range j.resMap["rows"].([]any) {
				row, err := dbhelper.MapToRow(v.(map[string]any))
				if err != nil {
					continue
				}
				dbhelper.CreateRow(db, row)
				// TODO: return the newly created rows
			}

		}

		j.jSONifyResMap()
		if j.err != nil {
			log.Fatal(j.err)

		}
		w.Write(j.resJSON)
	}
}

// takes in json with ["limit"] and ["offset"] keys
// returns json with ["rows"] key containing array of rows
// if any error occurs, ["errors"] key will contain array of error messages
func RetrieveHostsHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		setHeaderContentTypeJSON(&w)
		j := newJSONReqResVars()

		if r.Method != http.MethodPost {
			j.errors = append(j.errors, "Only POST method is allowed")
		}
		j.readAndParseReqJSON(r)

		if len(j.errors) == 0 {
			limit := 0
			offset := 0
			if j.resMap["limit"] != nil {
				limit = int(j.resMap["limit"].(float64))
			}
			if j.resMap["offset"] != nil {
				offset = int(j.resMap["offset"].(float64))
			}
			rows, err := dbhelper.RetrieveRows(db, limit, offset)
			if err != nil {
				j.errors = append(j.errors, err.Error())
			} else {
				j.resMap["rows"] = rows
			}
		}

		j.jSONifyResMap()
		w.Write(j.resJSON)
	}
}

// takes in json with ["rows"] key containing array of rows to update
// returns json with ["rows"] key containing array of updated rows with all attributes
// even if only one attribute is updated
// if any error occurs, ["errors"] key will contain array of error messages
func UpdateHostsHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		setHeaderContentTypeJSON(&w)
		j := newJSONReqResVars()

		if r.Method != http.MethodPost {
			j.errors = append(j.errors, "Only POST method is allowed")
		}
		j.readAndParseReqJSON(r)

		if len(j.errors) == 0 {
			updatedRows := make([]map[string]any, 0)
			for _, row := range j.resMap["rows"].([]any) {
				err := dbhelper.UpdateRow(db, row.(map[string]any))
				if err != nil {
					j.errors = append(j.errors, err.Error())
					continue
				}

				// Retrievig newly constructed row
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
			j.resMap["rows"] = updatedRows
		}

		// converting resMap to resJSON
		j.jSONifyResMap()
		w.Write(j.resJSON)
	}
}

// takes in json with ["rowIDs"] key containing array of rowIDs to delete
// returns json with ["deletedRowIDs"] key containing array of rowIDs that were successfully deleted
// if any error occurs, ["errors"] key will contain array of error messages
func DeleteHostsHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		setHeaderContentTypeJSON(&w)
		j := newJSONReqResVars()

		if r.Method != http.MethodPost {
			j.errors = append(j.errors, "Only Post method is allowed")
		}
		j.readAndParseReqJSON(r)

		deletedRowIDs := make([]int, 0)
		if len(j.errors) == 0 {
			for _, rawRowID := range j.resMap["rowIDs"].([]any) {
				rowID := int(rawRowID.(float64))
				success, err := dbhelper.DeleteRow(db, rowID)
				if err != nil {
					j.errors = append(j.errors, err.Error())
					continue
				}

				if success {
					deletedRowIDs = append(deletedRowIDs, rowID)
				}
			}
		}
		j.resMap["deletedRowIDs"] = deletedRowIDs

		j.jSONifyResMap()
		w.Write(j.resJSON)
	}
}
