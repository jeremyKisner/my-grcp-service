This is an example gRCP client-server project to learn.



## Generate service
Pre-requisite:
1. install [protoc](https://grpc.io/docs/protoc-installation/)

Run protoc executable like:
``` bash
D:\Programming\scripts\bin\protoc.exe --go_out=. --go-grpc_out=. .\rollerService\service.proto
```
