package internal_model

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
