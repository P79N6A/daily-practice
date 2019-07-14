#!/usr/bin/env bash

protoc -I IDL/ IDL/user.proto --go_out=plugins=grpc:services/gen
