#!/usr/bin/env bash

protoc -I IDL/ IDL/user.proto --go_out=plugins=grpc:services/grpc_gen

thrift -out services/thrift_gen -r --gen go IDL/user.thrift