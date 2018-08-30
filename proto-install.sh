wget "https://github.com/protocolbuffers/protobuf/releases/download/v3.6.1/protobuf-cpp-3.6.1.zip"

unzip protobuf-cpp-3.6.1.zip
cd protobuf-cpp-3.6.1
./autogen.sh

./configure --disable-shared

sudo make install

sudo ldconfig # refresh shared library cache.
# a new protoc will be located at
which protoc

go get -u github.com/golang/protobuf/protoc-gen-go
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
