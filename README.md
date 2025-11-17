# Go Networks

A pet project exploring network programming in Go. This project demonstrates various network concepts including TCP/UDP servers and clients, HTTP servers, and network utilities.

## Features

- **TCP Server/Client**: Basic TCP communication examples
- **UDP Server/Client**: UDP datagram examples
- **HTTP Server**: Simple HTTP server with routing
- **Network Utilities**: Port scanning, connection testing, and more

## Structure

```
.
├── cmd/
│   ├── tcp-server/     # TCP server implementation
│   ├── tcp-client/     # TCP client implementation
│   ├── udp-server/     # UDP server implementation
│   ├── udp-client/     # UDP client implementation
│   └── http-server/    # HTTP server implementation
├── pkg/
│   └── network/        # Network utility functions
└── examples/           # Example usage and demos
```

## Getting Started

### Prerequisites

- Go 1.21 or later

### Running Examples

```bash
# Run TCP server
go run cmd/tcp-server/main.go

# Run TCP client (in another terminal)
go run cmd/tcp-client/main.go

# Run UDP server
go run cmd/udp-server/main.go

# Run UDP client
go run cmd/udp-client/main.go

# Run HTTP server
go run cmd/http-server/main.go
```

## Examples

### TCP Server
```go
// Start a TCP server on port 8080
go run cmd/tcp-server/main.go
```

### TCP Client
```go
// Connect to TCP server
go run cmd/tcp-client/main.go localhost:8080
```

## License

MIT

