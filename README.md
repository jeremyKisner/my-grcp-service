# my-grcp-service

This is an example gRCP client-server project in Go.


## Generate Service

Pre-requisite:
1. install [protoc](https://grpc.io/docs/protoc-installation/)

Run protoc executable like:
``` bash
D:\Programming\scripts\bin\protoc.exe --go_out=. --go-grpc_out=. .\roller\rollerService\service.proto
```

## Run Tests

``` bash
go test ./...
```

## Build Locally

Build it!
``` bash
go build -o ./bin/roller_server ./roller/roller_server
```

## Locally Run Service with Docker

Build it!
``` bash
docker build -t roller_server .
```

Serve it!
``` bash
docker run -p 50051:50051 roller_server
```
