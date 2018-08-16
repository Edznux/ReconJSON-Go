package tests

import (
	"bufio"
	"net"
	"os"
	"testing"

	rjson "github.com/edznux/ReconJSON-Go"
)

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
	file, _ := os.Open(fileName)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		count++
	}

	if count != 5 {
		t.Errorf("The written file is not correctly wrote. Expected %d lines, got %d", 5, count)
	}
}

func TestMustJsonName(t *testing.T) {
	wrongFileName := "wrongname"
	hosts := []rjson.Host{rjson.Host{}}

	err := rjson.Write(hosts, wrongFileName)
	if err == nil {
		t.Errorf("Shouldn't be able to write file without .json extention, got %s", wrongFileName)
	}
}

func TestHost(t *testing.T) {
	var hosts []rjson.Host
	fileName := "tests.json"
	ip, _, _ := net.ParseCIDR("192.168.0.1")

	basicHost := rjson.Host{
		Type:    "Host",
		Fqdn:    "example.acme.com",
		IP:      ip,
		Domain:  "acme.com",
		Company: "Acme",
	}

	hosts = append(hosts, basicHost)
	err := rjson.Write(hosts, fileName)

	if err != nil {
		t.Errorf("Failed to write the ReconJSON file: %s", err.Error())
	}
}
