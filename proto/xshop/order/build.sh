#!/bin/bash
protoc --go_out=. *.proto -I. -I$GOPATH/src/github.com/my0sot1s/mpb/proto/xshop/
echo "done"