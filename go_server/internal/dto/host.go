package dto

import (
	"encoding/json"
	utils "themynet"
	"themynet/internal/model"
)

type HostDTO struct {
	Id              *int64
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
	CreatedBy       *string
	CreatedAt       *int64
	RecordedAt      *int64
}

type CreateHostsRequestDTO struct {
	Hosts []HostDTO
}

type RetrieveHostsRequestDTO struct {
	Limit  uint
	Offset uint
}

type UpdateHostsRequestDTO struct {
	Hosts []map[string]any // must contain id, checked in services
}

type DeleteHostsRequestDTO struct {
	HostIds []int64
}

type HostResponseDTO struct {
	Errors []string
	Hosts  []HostDTO
}

func (h *HostResponseDTO) AddErrors(errs ...string) {
	h.Errors = append(h.Errors, errs...)

}
func (h *HostResponseDTO) AddHostDTO(hostDTO HostDTO) {
	h.Hosts = append(h.Hosts, hostDTO)
}

func ModelToDTO(host model.Host) (hostDTO HostDTO, err error) {
	jsnHost, err := json.Marshal(host)
	if err != nil {
		return
	}

	err = json.Unmarshal(jsnHost, &hostDTO)

	return
}

func DtoToModel(hostDTO HostDTO) (host *model.Host, err error) {
	jsnHostDTO, err := json.Marshal(hostDTO)
	if err != nil {
		return
	}

	host = new(model.Host)
	err = json.Unmarshal(jsnHostDTO, host)
	utils.DebugPrintf("host: %+v\n", host)

	return
}

func MapToDTO(hostMap map[string]any) (hostDTO HostDTO, err error) {
	jsnHostMap, err := json.Marshal(hostMap)
	if err != nil {
		return
	}

	err = json.Unmarshal(jsnHostMap, &hostDTO)

	return
}
