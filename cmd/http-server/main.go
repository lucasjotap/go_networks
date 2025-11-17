package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
		<html>
		<head><title>Go Networks HTTP Server</title></head>
		<body>
			<h1>Welcome to Go Networks!</h1>
			<p>This is a simple HTTP server built with Go's net/http package.</p>
			<ul>
				<li><a href="/time">Current Time</a></li>
				<li><a href="/info">Server Info</a></li>
			</ul>
		</body>
		</html>
	`)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Current time: %s\n", time.Now().Format(time.RFC3339))
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
		<h2>Server Information</h2>
		<p><strong>Method:</strong> %s</p>
		<p><strong>URL:</strong> %s</p>
		<p><strong>Remote Address:</strong> %s</p>
		<p><strong>User Agent:</strong> %s</p>
	`, r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent())
}

func main() {
	port := ":8082"
	
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/time", timeHandler)
	http.HandleFunc("/info", infoHandler)
	
	fmt.Printf("HTTP Server listening on %s\n", port)
	fmt.Println("Available endpoints:")
	fmt.Println("  - http://localhost:8082/")
	fmt.Println("  - http://localhost:8082/time")
	fmt.Println("  - http://localhost:8082/info")
	
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

