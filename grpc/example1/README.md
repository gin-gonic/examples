# gRPC Example

First Step: run grpc server

```sh
go run grpc/server.go
```

Second Step: run gin server

```sh
go run gin/main.go
```

Testing command.

```sh
curl -v 'http://localhost:8052/rest/n/thinkerou'
```
