package main

import (
	"flag"
	"log"

	"grpcdemo/pb"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"io"
	"strconv"
	"os"
)

const serverHost = "localhost:9000"

func main() {
	option := flag.Int("o", 1, "Command to run")
	badgeNumber := flag.Int("b", 2080, "Badge number to get")
	filename := flag.String("f", "img.jpg", "File to send.")
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
	case 2:
		log.Println("Getting by badge number: ", *badgeNumber)
		getByBadgeNumber(client, *badgeNumber)
	case 3:
		log.Println("Getting all employees")
		getAll(client)
	case 4:
		log.Println("Adding Photo")
		addPhoto(client, *badgeNumber, *filename)
	}
}

func addPhoto(client pb.EmployeeServiceClient, badgeNumber int, filename string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal("Error opening the file: ", err)
	}
	defer f.Close()
	md := metadata.New(map[string]string{"badgenumber": strconv.Itoa(badgeNumber)})
	ctx := metadata.NewContext(context.Background(), md)
	stream, err := client.AddPhoto(ctx)
	if err != nil {
		log.Println(err)
	}

	for {
		chunk := make([]byte, 64*1024) // 64kb
		n, err := f.Read(chunk)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if n < len(chunk) {
			chunk = chunk[:n]
		}
		stream.Send(&pb.AddPhotoRequest{Data: chunk})
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Println(err)
	}
	log.Printf("Ok? %v", res.IsOk)
}

func getAll(client pb.EmployeeServiceClient) {
	stream, err := client.GetAll(context.Background(), &pb.GetAllRequest{})
	if err != nil {
		log.Fatal("Error getting all employees: ", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("error: ", err)
		}
		log.Printf("Got employee: %v\n", res.Employee)
	}
}

func getByBadgeNumber(client pb.EmployeeServiceClient, badgeNumber int) {
	res, err := client.GetByBadgeNumber(context.Background(), &pb.GetByBadgeNumberRequest{
		BadgeNumber: int32(badgeNumber),
	})
	if err != nil {
		log.Fatal("error: ", err)
	}
	log.Printf("got employee: %+v\n", res.Employee)
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
