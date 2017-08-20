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

# running mongo replica set
# replicate travel 1
mongod \
  --replSet travel \
  --dbpath ~/data/rep1 \
  --port 27001
# replicate travel 2
mongod \
  --replSet travel \
  --dbpath ~/data/rep2 \
  --port 27002
# replicate travel 3
mongod \
  --replSet travel \
  --dbpath ~/data/rep3 \
  --port 27003
# log into a node and initialize the replication set
mongo localhost:27001
rs.initiate({
  _id: 'travel',
  members: [
    {_id: 1, host: 'localhost:27001'},
    {_id: 2, host: 'localhost:27002'},
    {_id: 3, host: 'localhost:27003'}
  ]
})
# get the status of the replica set
rs.status()
rs.printSlaveReplicationInfo()
