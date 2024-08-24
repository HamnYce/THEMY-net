package apiv1handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"themynet/internal/dto"
	"themynet/internal/services"
)

func writeFatalError(w http.ResponseWriter, errS string) {
	w.Write([]byte(fmt.Sprintf("{\"error\": %q}", errS)))
}

func CreateHostsHandler(w http.ResponseWriter, r *http.Request) {
	createHostsRequestDTO := dto.CreateHostsRequestDTO{}

	err := json.NewDecoder(r.Body).Decode(&createHostsRequestDTO)
	if err != nil {
		writeFatalError(w, err.Error())
		return
	}

	if len(createHostsRequestDTO.Hosts) == 0 {
		writeFatalError(w, "Hosts not provided")
		return
	}

	// DATABASE INSERTION
	hostResponseDTO := services.CreateHosts(createHostsRequestDTO)
	fmt.Println("len(createHostsResponseDTO)", len(hostResponseDTO.Hosts))

	jsn, err := json.Marshal(hostResponseDTO)

	if err != nil {
		// NOTE: only some weird internal error can cause this
		writeFatalError(w, err.Error())
		return
	}

	w.Write(jsn)
}

func RetrieveHostsHandler(w http.ResponseWriter, r *http.Request) {
	retrieveHostsRequestDTO := dto.RetrieveHostsRequestDTO{}

	err := json.NewDecoder(r.Body).Decode(&retrieveHostsRequestDTO)
	if err != nil {
		writeFatalError(w, err.Error())
		return
	}

	// DATABASE RETRIEVAL
	hostResponseDTO := services.RetrieveHosts(retrieveHostsRequestDTO)

	jsn, err := json.Marshal(hostResponseDTO)
	if err != nil {
		// NOTE: only some weird internal error can cause this
		writeFatalError(w, err.Error())
		return
	}

	w.Write(jsn)
}

func UpdateHostsHandler(w http.ResponseWriter, r *http.Request) {
	updateHostsRequestDTO := dto.UpdateHostsRequestDTO{}

	err := json.NewDecoder(r.Body).Decode(&updateHostsRequestDTO)

	if err != nil {
		writeFatalError(w, err.Error())
		return
	}

	// DATABASE UPDATE
	hostResponseDTO := services.UpdateHosts(updateHostsRequestDTO)

	jsn, err := json.Marshal(hostResponseDTO)
	if err != nil {
		// NOTE: only some weird internal error can cause this
		writeFatalError(w, err.Error())
		return
	}

	w.Write(jsn)
}

func DeleteHostsHandler(w http.ResponseWriter, r *http.Request) {
	deleteHostsRequestDTO := dto.DeleteHostsRequestDTO{}

	err := json.NewDecoder(r.Body).Decode(&deleteHostsRequestDTO)
	if err != nil {
		writeFatalError(w, err.Error())
		return
	}

	// DATABASE DELETE
	hostResponseDTO := services.DeleteHosts(deleteHostsRequestDTO)

	jsn, err := json.Marshal(hostResponseDTO)
	if err != nil {
		// NOTE: only some weird internal error can cause this
		writeFatalError(w, err.Error())
		return
	}

	w.Write(jsn)
}
