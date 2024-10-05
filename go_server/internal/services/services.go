package services

import (
	"fmt"
	"strconv"
	"themynet/internal/db/sqls/sqls"
	"themynet/internal/dto"
	"themynet/internal/model"
)

func CreateHosts(createHostsRequestDTO dto.CreateHostsRequestDTO) (hostResponseDTO dto.HostResponseDTO) {
	// DATABASE INSERTION
	for _, hostDTO := range createHostsRequestDTO.Hosts {
		// DTO -> Model
		host, err := dto.DtoToModel(hostDTO)
		if err != nil {
			hostResponseDTO.AddErrors("Could not convert DTO to host: " + err.Error())
			continue
		}

		// DB
		createdHost, err := model.Createhost(sqls.CreateHostParams{})
		if err != nil {
			hostResponseDTO.AddErrors("Host could not be created: " + err.Error())
			continue
		}
		host = &createdHost

		// Model -> DTO
		createdHostDTO, err := dto.ModelToDTO(*host)
		if err != nil {
			hostResponseDTO.AddErrors("Host created and retrieved but could not parse into DTO, id=" + strconv.FormatInt(createdHost.ID.(int64), 10) + ": " + err.Error())
			continue
		}

		hostResponseDTO.Hosts = append(hostResponseDTO.Hosts, createdHostDTO)
	}

	return
}

func RetrieveHosts(retrieveHostsRequestDTO dto.RetrieveHostsRequestDTO) (hostResponseDTO dto.HostResponseDTO) {
	hosts, err := model.Retrievehosts(retrieveHostsRequestDTO.Limit, retrieveHostsRequestDTO.Offset)
	fmt.Printf("hosts: %+v", hosts)

	if err != nil {
		hostResponseDTO.AddErrors("could not retrieve hosts: " + err.Error())
		return
	}

	for _, host := range hosts {
		hostDTO, err := dto.ModelToDTO(host)
		if err != nil {
			hostResponseDTO.AddErrors("could not convert retrieved model to DTO, id=" + strconv.FormatInt(host.ID.(int64), 10) + ": " + err.Error())
			continue
		}

		hostResponseDTO.Hosts = append(hostResponseDTO.Hosts, hostDTO)
	}

	return
}

func UpdateHosts(updateHostsRequestDTO dto.UpdateHostsRequestDTO) (hostResponseDTO dto.HostResponseDTO) {
	var validHostMaps []map[string]any
	for _, hostMap := range updateHostsRequestDTO.Hosts {
		// ID CHECK
		if hostMap["Id"] == nil {
			hostResponseDTO.AddErrors("host did not contain Id, skipping")
			continue
		}

		// TYPE CHECKING: by converting to dto but discarding the result and only keeping the error
		// we discard the return value because of intentional nil values
		_, err := dto.MapToDTO(hostMap)
		if err != nil {
			hostResponseDTO.AddErrors(err.Error())
			continue
		}

		validHostMaps = append(validHostMaps, hostMap)
	}

	hosts, err := model.Updatehosts(validHostMaps)
	if err != nil {
		hostResponseDTO.AddErrors("hosts were all valid but db had error: " + err.Error())
		return
	}

	for _, host := range hosts {
		hostDTO, err := dto.ModelToDTO(host)
		if err != nil {
			hostResponseDTO.AddErrors("error converting host to hostDTO, id=" + string(host.ID.(int64)) + ": " + err.Error())
			continue
		}
		hostResponseDTO.AddHostDTO(hostDTO)
	}

	return
}

func DeleteHosts(deleteHostsRequestDTO dto.DeleteHostsRequestDTO) (hostResponseDTO dto.HostResponseDTO) {
	for _, hostId := range deleteHostsRequestDTO.HostIds {
		// we retrieve first to send back to the user the deleted models
		host, err := model.Retrievehost(hostId)
		if err != nil {
			hostResponseDTO.AddErrors("could not retrieve host with id=" + strconv.FormatInt(hostId, 10) + ": " + err.Error())
			continue
		}

		hostDTO, err := dto.ModelToDTO(*host)
		if err != nil {
			hostResponseDTO.AddErrors("could not convert host to dto with id=" + strconv.FormatInt(hostId, 10) + ": " + err.Error())
			continue
		}

		err = model.Deletehost(hostId)

		if err != nil {
			hostResponseDTO.AddErrors("db error while deleting host with id=" + strconv.FormatInt(hostId, 10) + ": " + err.Error())
			continue
		}

		hostResponseDTO.Hosts = append(hostResponseDTO.Hosts, hostDTO)
	}

	return
}
