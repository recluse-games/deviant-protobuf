#!/usr/bin/env bash

# This should only be run by a robot
PROTOC_ZIP=protoc-3.7.1-linux-x86_64.zip
curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.7.1/$PROTOC_ZIP
unzip -o $PROTOC_ZIP -d /usr/local bin/protoc
unzip -o $PROTOC_ZIP -d /usr/local 'include/*'
rm -f $PROTOC_ZIP

protoc --proto_path='protobuf/' --go_out=plugins=grpc:genproto/go/ protobuf/*
protoc --proto_path='protobuf/' --csharp_out=genproto/cs/ protobuf/*
