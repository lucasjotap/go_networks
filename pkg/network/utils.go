package network

import (
	"fmt"
	"net"
	"time"
)

// CheckPort checks if a port is open on the given host
func CheckPort(host string, port int, timeout time.Duration) (bool, error) {
	address := fmt.Sprintf("%s:%d", host, port)
	
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return false, nil
	}
	defer conn.Close()
	
	return true, nil
}

// ScanPorts scans a range of ports on a host
func ScanPorts(host string, startPort, endPort int, timeout time.Duration) []int {
	var openPorts []int
	
	for port := startPort; port <= endPort; port++ {
		if isOpen, _ := CheckPort(host, port, timeout); isOpen {
			openPorts = append(openPorts, port)
		}
	}
	
	return openPorts
}

// GetLocalIP returns the local IP address
func GetLocalIP() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()
	
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String(), nil
}

// IsValidIP checks if a string is a valid IP address
func IsValidIP(ip string) bool {
	return net.ParseIP(ip) != nil
}

// ResolveHostname resolves a hostname to IP addresses
func ResolveHostname(hostname string) ([]string, error) {
	ips, err := net.LookupIP(hostname)
	if err != nil {
		return nil, err
	}
	
	var ipStrings []string
	for _, ip := range ips {
		ipStrings = append(ipStrings, ip.String())
	}
	
	return ipStrings, nil
}

