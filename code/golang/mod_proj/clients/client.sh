#!/usr/bin/env bash

grpcurl -import-path ../IDL -proto ../IDL/user.proto \
 -plaintext -d '{"name":"moye"}' :1551 core_service.UserService/GetUser