#!/usr/bin/env bash

# Use for local development/manual invocation
protoc --proto_path='protobuf/' --go_out=plugins=grpc:genproto/go/ protobuf/*
protoc --proto_path='protobuf/' --csharp_out=genproto/cs/ protobuf/*
