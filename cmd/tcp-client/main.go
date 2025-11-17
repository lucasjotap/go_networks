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
		fmt.Println("Example: go run main.go localhost:8080")
		os.Exit(1)
	}
	
	address := os.Args[1]
	
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	
	fmt.Printf("Connected to %s\n", address)
	fmt.Println("Type messages (or 'quit' to exit):")
	
	// Read responses in a goroutine
	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			log.Printf("Error reading: %v", err)
		}
	}()
	
	// Read input from stdin
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		
		if text == "quit" {
			conn.Write([]byte("quit\n"))
			return
		}
		
		if text != "" {
			conn.Write([]byte(text + "\n"))
		}
	}
}

