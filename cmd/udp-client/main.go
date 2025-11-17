package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <host:port>")
		fmt.Println("Example: go run main.go localhost:8081")
		os.Exit(1)
	}
	
	address := os.Args[1]
	
	serverAddr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		log.Fatalf("Failed to resolve address: %v", err)
	}
	
	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	
	fmt.Printf("Connected to UDP server at %s\n", address)
	fmt.Println("Type messages (or 'quit' to exit):")
	
	// Read responses in a goroutine
	go func() {
		buffer := make([]byte, 1024)
		for {
			n, err := conn.Read(buffer)
			if err != nil {
				log.Printf("Error reading: %v", err)
				return
			}
			fmt.Println(string(buffer[:n]))
		}
	}()
	
	// Read input from stdin
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		
		if text == "quit" {
			return
		}
		
		if text != "" {
			_, err := conn.Write([]byte(text))
			if err != nil {
				log.Printf("Error writing: %v", err)
			}
		}
	}
}

