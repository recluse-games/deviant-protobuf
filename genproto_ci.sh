#!/usr/bin/env bash

# This should only be run by a robot
PB_REL="https://github.com/protocolbuffers/protobuf/releases"
sudo curl -LO $PB_REL/download/v3.11.4/protoc-3.11.4-linux-x86_64.zip
sudo unzip protoc-3.11.4-linux-x86_64.zip -d $HOME/.local
export PATH="$PATH:$HOME/.local/bin"
go get github.com/golang/protobuf/protoc-gen-go
export PATH="$PATH:$(go env GOPATH)/bin"

protoc --proto_path='protobuf/' --go_out=plugins=grpc:genproto/go/ protobuf/*
protoc --proto_path='protobuf/' --csharp_out=genproto/cs/ protobuf/*
