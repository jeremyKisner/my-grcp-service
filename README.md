# my-grcp-service
This is an example gRCP client-server project in Go.


## Generate Service

Pre-requisite:
1. install [protoc](https://grpc.io/docs/protoc-installation/)

Run protoc executable like:
``` bash
D:\Programming\scripts\bin\protoc.exe --go_out=. --go-grpc_out=. .\rollerService\service.proto
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
