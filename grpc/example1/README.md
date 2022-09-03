# gRPC Example

This guide gets you started with gRPC in Go with a simple working example.

## Prerequisites

Install the protocol compiler plugins for Go using the following commands:

```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

Update your `PATH` so that the `protoc` compiler can find the plugins:

```sh
export PATH="$PATH:$(go env GOPATH)/bin"
```

## Regenerate gRPC code

```sh
protoc --go_out=gen --go_opt=paths=source_relative \
  --go-grpc_out=gen --go-grpc_opt=paths=source_relative \
  -I=$PWD pb/helloworld.proto
```

## Runing

First Step: run grpc server

```sh
go run grpc/server.go
```

Second Step: run gin server

```sh
go run gin/main.go
```

## Testing

Send data to gin server:

```sh
curl -v 'http://localhost:8080/rest/n/gin'
```

or using [grpcurl](https://github.com/fullstorydev/grpcurl) command:

```sh
grpcurl -d '{"name": "gin"}' \
  -plaintext localhost:50051 helloworld.v1.Greeter/SayHello
```
