#!/usr/bin/env bash

# Setup Windows Protoc Version
curl -L https://www.nuget.org/api/v2/package/Grpc.Tools/ -o /tmp/grpctools.zip 
unzip -o /tmp/grpctools.zip -d /tmp/grpctools/
chmod +x /tmp/grpctools/tools/linux_x64/protoc
chmod +x /tmp/grpctools/tools/linux_x64/grpc_csharp_plugin

# Use for local development/manual invocation
protoc --proto_path='protobuf/packages/directory/' --go_out=plugins=grpc:genproto/go/directory/ protobuf/packages/directory/*.proto
/tmp/grpctools/tools/linux_x64/protoc --plugin=protoc-gen-grpc=/tmp/grpctools/tools/linux_x64/grpc_csharp_plugin --proto_path='protobuf/packages/directory/' --csharp_out=genproto/cs/directory --grpc_out=genproto/cs/directory/  protobuf/packages/directory/*.proto

protoc --proto_path='protobuf/packages/instance_shard/' --go_out=plugins=grpc:genproto/go/instance_shard protobuf/packages/instance_shard/*.proto
/tmp/grpctools/tools/linux_x64/protoc --plugin=protoc-gen-grpc=/tmp/grpctools/tools/linux_x64/grpc_csharp_plugin --proto_path='protobuf/packages/instance_shard/' --csharp_out=genproto/cs/instance_shard --grpc_out=genproto/cs/instance_shard/  protobuf/packages/instance_shard/*.proto