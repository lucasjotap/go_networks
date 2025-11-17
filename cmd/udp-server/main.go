package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	port := ":8081"
	
	addr, err := net.ResolveUDPAddr("udp", port)
	if err != nil {
		log.Fatalf("Failed to resolve address: %v", err)
	}
	
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	defer conn.Close()
	
	fmt.Printf("UDP Server listening on %s\n", port)
	fmt.Println("Waiting for messages...")
	
	buffer := make([]byte, 1024)
	
	for {
		n, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Printf("Error reading: %v", err)
			continue
		}
		
		message := strings.TrimSpace(string(buffer[:n]))
		fmt.Printf("Received from %s: %s\n", clientAddr, message)
		
		// Echo back the message
		response := fmt.Sprintf("Echo: %s", message)
		conn.WriteToUDP([]byte(response), clientAddr)
	}
}

