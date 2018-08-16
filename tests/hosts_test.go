package tests

import (
	"bufio"
	"errors"
	"net"
	"os"
	"testing"

	rjson "github.com/edznux/ReconJSON-Go"
)

func RemoveFile(filename string) error {
	err := os.Remove(filename)
	if err != nil {
		return errors.New("Could not delete file : " + err.Error())
	}
	return nil
}
func TestMustOneByLine(t *testing.T) {
	fileName := "testLine.json"

	var hosts []rjson.Host
	ip, _, _ := net.ParseCIDR("192.168.0.1")

	basicHost := rjson.Host{
		Type:    "Host",
		Fqdn:    "example.acme.com",
		IP:      ip,
		Domain:  "acme.com",
		Company: "Acme",
	}

	// should add 3 line
	hosts = append(hosts, basicHost)
	hosts = append(hosts, basicHost)
	hosts = append(hosts, basicHost)

	err := rjson.Write(hosts, fileName)

	if err != nil {
		t.Errorf("Failed to write the ReconJSON file: %s", err.Error())
	}

	count := 0
	file, err := os.Open(fileName)
	if err != nil {
		t.Errorf("Failed to open file %s : %s", fileName, err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		count++
	}

	if count != 5 {
		t.Errorf("The written file is not correctly wrote. Expected %d lines, got %d", 5, count)
	}

	RemoveFile(fileName)
}

func TestMustJsonName(t *testing.T) {
	wrongFileName := "wrongname"
	hosts := []rjson.Host{rjson.Host{}}

	err := rjson.Write(hosts, wrongFileName)
	if err == nil {
		t.Errorf("Shouldn't be able to write file without .json extention, got %s", wrongFileName)
		RemoveFile(wrongFileName)
	}
}

func TestHost(t *testing.T) {
	var hosts []rjson.Host
	fileName := "testHost.json"
	ip, _, _ := net.ParseCIDR("192.168.0.1")

	basicHost := rjson.NewHost()
	basicHost.Type = "Host"
	basicHost.Fqdn = "example.acme.com"
	basicHost.IP = ip
	basicHost.Domain = "acme.com"
	basicHost.Company = "Acme"

	hosts = append(hosts, *basicHost)

	err := rjson.Write(hosts, fileName)
	if err != nil {
		t.Errorf("Failed to write the ReconJSON file: %s", err.Error())
	}

	RemoveFile(fileName)
}

func TestHostPort(t *testing.T) {
	var hosts []rjson.Host
	fileName := "testPort.json"
	ip, _, _ := net.ParseCIDR("192.168.0.1")
	port := rjson.NewPort()
	port.State = rjson.StateOpen.String()
	port.Port = 22
	port.Protocol = "ssh"
	port.Banner = "SSH-2.0-OpenSSH_7.2p2 Ubuntu-4ubuntu2.4"

	TCPPorts := []rjson.Port{}
	TCPPorts = append(TCPPorts, *port)

	basicHost := rjson.NewHost()

	basicHost.Type = "Host"
	basicHost.Fqdn = "example.acme.com"
	basicHost.IP = ip
	basicHost.Domain = "acme.com"
	basicHost.Company = "Acme"
	basicHost.Ports.TCP = TCPPorts

	hosts = append(hosts, *basicHost)
	err := rjson.Write(hosts, fileName)

	if err != nil {
		t.Errorf("Failed to write the ReconJSON file: %s", err.Error())
	}
	RemoveFile(fileName)
}

func TestHostDNS(t *testing.T) {
	var hosts []rjson.Host
	fileName := "testDNS.json"
	ip, _, _ := net.ParseCIDR("192.168.0.1")

	dns := rjson.NewDNS()
	dns.A = []string{"192.168.0.1", "192.168.0.2"}
	dns.AAAA = []string{"fe80::1"}
	dns.CNAME = []string{"ex.acme.com"}
	dns.PTR = []string{"ex.acme.com"}
	dns.MX = []string{"example-acme-com.mail.protection.outlook.com"}
	dns.NS = []string{"nameserver.acme.com"}
	dns.TXT = []string{"txtRecordString"}

	basicHost := rjson.NewHost()

	basicHost.Fqdn = "example.acme.com"
	basicHost.IP = ip
	basicHost.Domain = "acme.com"
	basicHost.Company = "Acme"
	basicHost.DNS = *dns

	hosts = append(hosts, *basicHost)
	err := rjson.Write(hosts, fileName)

	if err != nil {
		t.Errorf("Failed to write the ReconJSON file: %s", err.Error())
	}
	RemoveFile(fileName)
}

func TestHostService(t *testing.T) {
	var hosts []rjson.Host
	fileName := "testService.json"
	ip, _, err := net.ParseCIDR("192.168.0.1/24")

	if err != nil {
		t.Errorf("Could parse CIDR IP")
	}

	service := rjson.NewService()
	service.Protocol = "http"
	service.Content = map[string]string{
		"path":         "/test",
		"screenshot":   "/root/screenshots/screenshot.jpg",
		"code":         "200",
		"content-type": "text/html",
		"length":       "1024",
	}

	port := rjson.NewPort()
	port.State = rjson.StateOpen.String()
	port.Port = 80
	port.Protocol = "http"
	port.Service = *service

	TCPPorts := []rjson.Port{}
	TCPPorts = append(TCPPorts, *port)

	basicHost := rjson.NewHost()
	basicHost.Type = "Host"
	basicHost.Fqdn = "example.acme.com"
	basicHost.IP = ip
	basicHost.Domain = "acme.com"
	basicHost.Company = "Acme"
	basicHost.Ports.TCP = TCPPorts

	hosts = append(hosts, *basicHost)
	err = rjson.Write(hosts, fileName)

	if err != nil {
		t.Errorf("Failed to write the ReconJSON file: %s", err.Error())
	}

	RemoveFile(fileName)

}
