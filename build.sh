#!/bin/bash
# build for product
protoc -I product/  product/product.proto  --go_out=plugins=grpc:product
# build for user
protoc -I user/ user/user.proto --go_out=./user
#build for req
protoc -I reqres/ reqres/reqres.proto --go_out=./reqres
# build for blog
protoc -I blog/ blog/blog.proto --go_out=./blog
echo done!