package reconjson

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"os"
	"strings"
)

type PortState uint8

const (
	Closed = PortState(iota)
	Open
	Filtered
)

func (p PortState) String() string {
	name := []string{"closed", "open", "filtered"}
	i := uint8(p)

	switch {
	case i <= uint8(Filtered):
		return name[i]
	default:
		// this NEVER should be the case. Panic seems the correct behavior
		panic("This type of port state doesn't exist")
	}
}

type DNS struct {
	Type  string   `json:"type,omitempty"` // will be "DNS"
	A     []string `json:"A,omitempty"`
	AAAA  []string `json:"AAAA,omitempty"`
	CNAME []string `json:"CNAME,omitempty"`
	PTR   []string `json:"PTR,omitempty"`
	MX    []string `json:"MX,omitempty"`
	TXT   []string `json:"TXT,omitempty"`
}

type Service struct {
	Type     string            `json:"type,omitempty"` // will be "Service"
	Protocol string            `json:"protocol"`
	Content  map[string]string `json:"content,omitempty"`
}

type Port struct {
	Type     string  `json:"type,omitempty"` // will be "Port"
	Port     uint16  `json:"port"`
	State    string  `json:"state"`
	Protocol string  `json:"protocol,omitempty"`
	Banner   string  `json:"banner,omitempty"`
	Service  Service `json:"service,omitempty"`
}

type Host struct {
	Type    string `json:"type,omitempty"` // will be "Host"
	Fqdn    string `json:"fqdn,omitempty"`
	IP      net.IP `json:"ip,string"`
	Domain  string `json:"domain,omitempty"`
	Company string `json:"company,omitempty"`
	DNS     DNS    `json:"dns,omitempty"`
	Ports   struct {
		TCP []Port `json:"tcp,omitempty"`
		UDP []Port `json:"udp,omitempty"`
	} `json:"ports,omitempty"`
}

func NewHost() *Host {
	return &Host{Type: "host"}
}

func NewDNS() *DNS {
	return &DNS{Type: "dns"}
}

func NewService() *Service {
	return &Service{Type: "service"}
}

func NewPort() *Port {
	return &Port{Type: "port"}
}

/*
Write takes a slice of hosts and write it down into "filename"
It (tries to) follow the Recon.JSON data structure found here : https://github.com/ReconJSON/ReconJSON
One host by line, (execept opening and closing lines)
*/
func Write(hosts []Host, filename string) error {
	var lines []string

	if !strings.HasSuffix(filename, ".json") {
		return errors.New("The provided filename doesn't end with .json")
	}

	for _, h := range hosts {
		line, err := json.Marshal(h)
		if err != nil {
			return errors.New("Could not write json file : " + err.Error())
		}
		fmt.Println("line :", string(line))
		fmt.Println("h :", h)
		lines = append(lines, string(line))
	}
	data := "[\n"
	data += strings.Join(lines, ",\n")
	data += "\n]"

	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	_, err = f.WriteString(data)
	if err != nil {
		return err
	}
	f.Sync()

	return nil
}
