package main

import (
	"log"

	"grpcdemo/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const serverHost = "localhos:9000"

func main() {
	creds, err := credentials.NewClientTLSFromFile("cert.pem", "")
	if err != nil {
		log.Fatal(err)
	}

	opts := []grpc.DialOption{grpc.WithTransportCredentials(creds)}
	conn, err := grpc.Dial(serverHost, opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	pb.NewEmployeeServiceClient(conn)
}
