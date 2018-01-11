# Go-gRPC-examples
A simple implementation of GRPC using Golang

Prerequisites
============

###### Go

- Install [Go] (https://golang.org/dl/)
- Create Go workspace as mentioned [ here ] (https://golang.org/dl/)
- Set `GO_PATH` according to newly created workspace


###### Protobuf and gRPC

- Install protoc from [here] (https://github.com/google/protobuf/releases/tag/v3.3.0)
- Set `PROTOC_PATH`
- Install Go protobuf plugin `go get -a github.com/golang/protobuf/protoc-gen-go`

###### Dep 
Install package management tool dep for Golang

- go get -u github.com/golang/dep/cmd/dep
- RUN dep init (this command initializes the directory)

Building project
================

###### Get project
- Clone project to `GO_PATH` like `GO_PATH/src/Go-gRPC-examples`
- RUN dep ensure ( this command gets all dependencies)

###### Compiling proto files
- Go to the proto file directory `GO_PATH/src/proto/`
- Execute `protoc --go_out=plugins=grpc:. helloworld.proto`

###### Run Server 
- Go to directory Server
- Execute go run main.go

###### Run Client
- Go to directory Client
- Execute go run main.go



