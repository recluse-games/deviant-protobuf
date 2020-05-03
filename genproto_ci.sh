#!/usr/bin/env bash

# This should only be run by a robot
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOROOT:$GOPATH:$GOBIN

PROTOC_ZIP=protoc-3.7.1-linux-x86_64.zip
curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.7.1/$PROTOC_ZIP
sudo unzip -o $PROTOC_ZIP -d /usr/local bin/protoc
sudo unzip -o $PROTOC_ZIP -d /usr/local 'include/*'
rm -f $PROTOC_ZIP
go install google.golang.org/protobuf/cmd/protoc-gen-go

protoc --proto_path='protobuf/' --go_out=plugins=grpc:genproto/go/ protobuf/*
protoc --proto_path='protobuf/' --csharp_out=genproto/cs/ protobuf/*
