package main

import (
	"flag"
	"log"

	"grpcdemo/pb"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

const serverHost = "localhost:9000"

func main() {
	option := flag.Int("o", 1, "Command to run")
	flag.Parse()

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

	client := pb.NewEmployeeServiceClient(conn)
	switch *option {
	case 1:
		log.Println("sending metadata")
		sendMetadata(client)
	}
}

func sendMetadata(client pb.EmployeeServiceClient) {
	md := metadata.MD{}
	md["user"] = []string{"petergriffin"}
	md["email"] = []string{"peterg@msn.com"}
	md["password"] = []string{"pAsswOrdOnE"}

	ctx := metadata.NewContext(context.Background(), md)
	res, err := client.GetByBadgeNumber(ctx, &pb.GetByBadgeNumberRequest{})
	log.Println("Error", err)
	log.Printf("Response %+v\n", res)
}
