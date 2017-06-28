# install gRPC tooling
go get -u google.golang.org/grpc

# install protocol buffer tooling
go get -u github.com/golang/protobuf/proto
go get -u github.com/golang/protobuf/protc-gen-go

# compile proto files to go
protoc \
  -I ./pb \ # where the messages are
  ./pb/messages.proto \ # the exact message file
  --go_out=plugins=grpc:./src

# generate self signed cert and key
openssl req -x509 -nodes -days 3650 -newkey rsa:2048 -keyout key.pem -out cert.pem