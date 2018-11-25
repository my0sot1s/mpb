#!/bin/bash
protoc -I product/  product/product.proto  --go_out=plugins=grpc:product
protoc -I user/ user/user.proto --go_out=./user