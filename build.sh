#!/bin/bash
# # build for user
protoc user/user.proto -I. --go_out=plugins=grpc:.

protoc wsh/wsh.proto -I. --go_out=plugins=grpc:.

protoc auth/auth.proto -I. --go_out=plugins=grpc:.

echo done!