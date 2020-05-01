#!/bin/sh

protoc --proto_path='protobuf/' --go_out=plugins=grpc:genproto/ protobuf/*
