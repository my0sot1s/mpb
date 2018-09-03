#!/bin/bash
protoc --go_out=. *.proto -I. -I$GOPATH/src/
echo "done"