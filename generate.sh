#!/bin/bash
cd proto
protoc --go_out=./ *.proto -I. -I$GOPATH/src/
mv *.pb.go ../models/
cd ..

