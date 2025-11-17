package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/lucasjotap/go_networks/pkg/network"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <host> [start-port] [end-port]")
		fmt.Println("Example: go run main.go localhost 8000 8100")
		os.Exit(1)
	}

	host := os.Args[1]
	startPort := 8000
	endPort := 8010

	if len(os.Args) >= 3 {
		fmt.Sscanf(os.Args[2], "%d", &startPort)
	}
	if len(os.Args) >= 4 {
		fmt.Sscanf(os.Args[3], "%d", &endPort)
	}

	fmt.Printf("Scanning ports %d-%d on %s...\n", startPort, endPort, host)

	timeout := 500 * time.Millisecond
	openPorts := network.ScanPorts(host, startPort, endPort, timeout)

	if len(openPorts) == 0 {
		fmt.Println("No open ports found.")
	} else {
		fmt.Printf("Open ports: %v\n", openPorts)
	}
}

