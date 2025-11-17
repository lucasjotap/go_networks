package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	
	// Set read deadline
	conn.SetReadDeadline(time.Now().Add(30 * time.Second))
	
	// Send welcome message
	conn.Write([]byte("Welcome to TCP Server! Type 'quit' to exit.\n"))
	
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		
		if text == "quit" {
			conn.Write([]byte("Goodbye!\n"))
			return
		}
		
		if text == "" {
			continue
		}
		
		// Echo back the message
		response := fmt.Sprintf("Echo: %s\n", text)
		conn.Write([]byte(response))
	}
	
	if err := scanner.Err(); err != nil {
		log.Printf("Error reading from connection: %v", err)
	}
}

func main() {
	port := ":8080"
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	defer listener.Close()
	
	fmt.Printf("TCP Server listening on %s\n", port)
	fmt.Println("Waiting for connections...")
	
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}
		
		fmt.Printf("New connection from %s\n", conn.RemoteAddr())
		go handleConnection(conn)
	}
}

