package apiv1handlers

import (
	"encoding/json"
	"net/http"
	debug "themynet/internal/debug"
	model "themynet/internal/model"
)

// struct to hold variables used in all CRUD handlers

// START: jSONReqResVars ---------
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

// END: jSONReqResVars ---------

func setHeaderContentTypeJSON(w *http.ResponseWriter) {
	(*w).Header().Del("Content-Type")
	(*w).Header().Add("Content-Type", "application/json")
}

func CreateHostsHandler(w http.ResponseWriter, r *http.Request) {
	setHeaderContentTypeJSON(&w)
	j := newJSONReqResVars()

	if r.Method != http.MethodPost {
		j.errors = append(j.errors, "Only POST method is allowed")
	}
	j.readAndParseReqJSON(r)

	if j.resMap["hosts"] == nil {
		j.errors = append(j.errors, "Hosts not provided")
	}

	// TODO: change this so that the code reads something like
	// 	if len == 0 : services.createHosts()
	// DATABASE INSERTION
	if len(j.errors) == 0 {
		for _, v := range j.resMap["hosts"].([]any) {
			host := new(model.Host)
			err := host.ParseMap(v.(map[string]any))
			if err != nil {
				continue
			}
			host.Createhost()
		}
	}

	j.jSONifyResMap()
	debug.CheckAndFatal(j.err)

	w.Write(j.resJSON)
}

func RetrieveHostsHandler(w http.ResponseWriter, r *http.Request) {
	setHeaderContentTypeJSON(&w)
	j := newJSONReqResVars()

	if r.Method != http.MethodPost {
		j.errors = append(j.errors, "Only POST method is allowed")
	}

	j.readAndParseReqJSON(r)

	if j.resMap["limit"] == nil {
		j.errors = append(j.errors, "Limit not provided")
	}

	if j.resMap["offset"] == nil {
		j.errors = append(j.errors, "Offset not provided")
	}

	// DATABASE RETRIEVAL
	if len(j.errors) == 0 {
		limit := int(j.resMap["limit"].(float64))
		offset := int(j.resMap["offset"].(float64))
		hosts, err := model.Retrievehosts(limit, offset)
		if err != nil {
			j.errors = append(j.errors, err.Error())
		} else {
			j.resMap["hosts"] = hosts
		}
	}

	delete(j.resMap, "limit")
	delete(j.resMap, "offset")

	j.jSONifyResMap()
	debug.CheckAndFatal(j.err)
	w.Write(j.resJSON)

}

func UpdateHostsHandler(w http.ResponseWriter, r *http.Request) {
	setHeaderContentTypeJSON(&w)
	j := newJSONReqResVars()

	if r.Method != http.MethodPost {
		j.errors = append(j.errors, "Only POST method is allowed")
	}
	j.readAndParseReqJSON(r)

	// DATABASE UPDATE
	if len(j.errors) == 0 {
		updatedhosts := make([]model.Host, 0)
		for _, hostMap := range j.resMap["hosts"].([]any) {
			host := new(model.Host)
			err := host.Updatehost(hostMap.(map[string]any))
			if err != nil {
				j.errors = append(j.errors, err.Error())
				continue
			}

			// Retrievig newly constructed host

			updatedhost, err := model.Retrievehost(*host.Id)
			if err != nil {
				continue
			}

			updatedhosts = append(updatedhosts, updatedhost)
		}
		j.resMap["hosts"] = updatedhosts
	}

	// converting resMap to resJSON
	j.jSONifyResMap()
	debug.CheckAndFatal(j.err)
	w.Write(j.resJSON)

}

func DeleteHostsHandler(w http.ResponseWriter, r *http.Request) {
	setHeaderContentTypeJSON(&w)
	j := newJSONReqResVars()

	if r.Method != http.MethodPost {
		j.errors = append(j.errors, "Only Post method is allowed")
	}
	j.readAndParseReqJSON(r)

	// DATABASE DELETE
	deletedhostIDs := make([]int, 0)
	if len(j.errors) == 0 {
		for _, rawhostID := range j.resMap["hostIDs"].([]any) {
			host := new(model.Host)
			host.Id = new(int)
			*(host.Id) = int(rawhostID.(float64))

			success, err := host.Deletehost()
			if err != nil {
				j.errors = append(j.errors, err.Error())
				continue
			}

			if success {
				deletedhostIDs = append(deletedhostIDs, *host.Id)
			}
		}
	}
	j.resMap["deletedHostIDs"] = deletedhostIDs
	delete(j.resMap, "hostIDs")

	j.jSONifyResMap()
	debug.CheckAndFatal(j.err)
	w.Write(j.resJSON)
}
