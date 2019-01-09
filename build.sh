#!/bin/bash
# #build for common
protoc common/common.proto -I. --go_out=.
# build for product
protoc product/product.proto -I. --go_out=plugins=grpc:.
# # build for user
protoc user/user.proto -I. -I$GOPATH/src/github.com/my0sot1s/mpb --go_out=plugins=grpc:.
# # build for blog
protoc blog/blog.proto -I. --go_out=.
# # build for maker
protoc marker/marker.proto -I. --go_out=plugins=grpc:.
echo done!