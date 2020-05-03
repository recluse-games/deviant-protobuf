#!/bin/sh

protoc --proto_path='protobuf/' --go_out=plugins=grpc:genproto/go/ protobuf/*
protoc --proto_path='protobuf/' --csharp_out=genproto/cs/ protobuf/*
