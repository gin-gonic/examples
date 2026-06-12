
install: 
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

generate: clean
	protoc --go_out=. \
	--go-grpc_out=. \
	-I=$(PWD) pb/helloworld/v1/*.proto

clean:
	rm -rf gen
