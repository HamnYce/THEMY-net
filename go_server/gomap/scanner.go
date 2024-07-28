package gomap

import (
	"bytes"
	"encoding/json"

	"github.com/Ullaakut/nmap"
)

func NewScanner(addresses []string) (*nmap.Scanner, error) {
	return nmap.NewScanner(
		nmap.WithTargets(addresses...),
		nmap.WithFastMode(),
		nmap.WithOSDetection(),
		nmap.WithMaxParallelism(8),
		nmap.WithVerbosity(1),
	)
}

func ScanHosts(scanner *nmap.Scanner) (hosts []nmap.Host, jsn []byte, warnings []string, err error) {
	res, warnings, err := scanner.Run()

	if err != nil {
		return
	}

	outputBytes := make([][]byte, 0)
	outputBytes = append(outputBytes, []byte("["))
	for _, host := range res.Hosts {
		if len(host.Addresses) == 0 || len(host.Ports) == 0 {
			continue
		}

		hosts = append(hosts, host)

		hostJson, err := json.Marshal(host)

		if err != nil {
			return hosts, jsn, warnings, err
		}

		outputBytes = append(outputBytes, hostJson)
		outputBytes = append(outputBytes, []byte(","))
	}
	outputBytes = append(outputBytes, []byte("]"))
	jsn = bytes.Join(outputBytes, []byte(""))

	return
}
