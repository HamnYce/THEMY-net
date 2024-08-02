package gomap

import (
	"strings"

	"github.com/Ullaakut/nmap"
)

type Host struct {
	IPAddress  string
	MACAddress string
	Hostname   string
	HostType   string
	OsName     string
	OsVersion  string
	Ports      []uint16
	RecordedAt string
}

type Port struct {
	ID      uint16
	service string
	state   string
	reason  string
}

func NewHostScaner(addresses []string, mostCommonPorts int) (*nmap.Scanner, error) {
	return nmap.NewScanner(
		nmap.WithTargets(addresses...),
		nmap.WithMostCommonPorts(mostCommonPorts),
		nmap.WithReason(),
		nmap.WithOSDetection(),
		nmap.WithMaxParallelism(4),
	)
}

func ScanHosts(addresses []string, mostCommonPorts int) (hosts []Host, warnings []string, err error) {
	scanner, err := NewHostScaner(addresses, mostCommonPorts)
	if err != nil {
		return
	}

	res, warnings, err := scanner.Run()

	if err != nil {
		return
	}

	for _, nmapHost := range res.Hosts {
		if len(nmapHost.Addresses) == 0 || len(nmapHost.Ports) == 0 {
			continue
		}

		host := Host{}

		// formatted as a unix timestamp
		host.RecordedAt = nmapHost.StartTime.FormatTime()

		for _, address := range nmapHost.Addresses {
			if address.AddrType == "ipv4" {
				host.IPAddress = address.Addr
			} else if address.AddrType == "mac" {
				host.MACAddress = address.Addr
			}
		}

		if len(nmapHost.Hostnames) > 0 {
			host.Hostname = nmapHost.Hostnames[0].Name
			host.HostType = nmapHost.Hostnames[0].Type
		}

		for _, port := range nmapHost.Ports {
			host.Ports = append(host.Ports, port.ID)
		}

		if len(nmapHost.OS.Matches) > 0 {
			osInfo := nmapHost.OS.Matches[0].Name
			host.OsName = osInfo[:strings.Index(osInfo, " ")]
			host.OsVersion = osInfo[strings.Index(osInfo, " ")+1:]
		}

		hosts = append(hosts, host)
	}

	return
}

func hostsToCSV(hosts []Host) []byte {
	return []byte{}
}

func hostsToJSON(hosts []Host) []byte {
	return []byte{}
}
