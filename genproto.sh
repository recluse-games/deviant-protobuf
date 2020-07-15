#!/usr/bin/env bash

# Use for local development/manual invocation
protoc --proto_path='protobuf/packages/directory/' --go_out=plugins=grpc:genproto/go/directory/ protobuf/packages/directory/*.proto
protoc --proto_path='protobuf/packages/directory/' --csharp_out=genproto/cs/directory protobuf/packages/directory/*.proto

curl -L https://www.nuget.org/api/v2/package/Grpc.Tools/ -o /tmp/grpctools.zip
unzip /tmp/grpctools.zip

protoc --proto_path='protobuf/packages/instance_shard/' --go_out=plugins=grpc:genproto/go/instance_shard protobuf/packages/instance_shard/*.proto
protoc --proto_path='protobuf/packages/instance_shard/' --csharp_out=plugins=grpc:genproto/cs/instance_shard protobuf/packages/instance_shard/*.proto