# my-grcp-service

This is an example gRCP client-server project in Go.

## Getting Started
1. Install [protoc](https://grpc.io/docs/protoc-installation/)
2. Run protoc executable like:
``` bash
protoc.exe --go_out=. --go-grpc_out=. .\roller\rollerService\service.proto
```
3. Run unit tests
``` bash
go test ./...
```

## Build Locally
1. Start server.
``` bash
go run .\roller\roller_server\main.go
```
2. Start client.
``` bash
go run .\roller\roller_client\main.go
```
3. While the server is running, use the client to send commands. The following keybinds exist:
    - **r** :: rolls a number 0-100 against the server.

## Locally Run Service with Docker
1. Build the server.
``` bash
docker build -t roller_server .
```
2. Start server.
``` bash
docker run -p 50051:50051 roller_server
```
3. Start client.
``` bash
go run .\roller\roller_client\main.go
```
