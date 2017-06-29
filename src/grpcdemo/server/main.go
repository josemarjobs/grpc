package main

import (
	"errors"
	"fmt"
	"grpcdemo/pb"
	"log"
	"net"

	"google.golang.org/grpc/credentials"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"io"
)

const port = ":9000"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	creds, err := credentials.NewServerTLSFromFile("cert.pem", "key.pem")
	if err != nil {
		log.Fatal(err)
	}

	opts := []grpc.ServerOption{grpc.Creds(creds)}
	s := grpc.NewServer(opts...)
	pb.RegisterEmployeeServiceServer(s, new(employeeService))
	log.Println("starting server on port: " + port)
	s.Serve(lis)
}

type employeeService struct {
}

func (s *employeeService) GetByBadgeNumber(ctx context.Context,
	req *pb.GetByBadgeNumberRequest) (*pb.EmployeeResponse, error) {

	if md, ok := metadata.FromContext(ctx); ok {
		fmt.Printf("Metadata: %+v\n", md)
	}
	for _, e := range employees {
		if e.BadgeNumber == req.BadgeNumber {
			return &pb.EmployeeResponse{Employee: &e}, nil
		}
	}
	return nil, errors.New("employee not found")
}

func (s *employeeService) GetAll(req *pb.GetAllRequest,
	stream pb.EmployeeService_GetAllServer) error {
	for _, e := range employees{
		stream.Send(&pb.EmployeeResponse{Employee: &e})
	}

	return nil
}

func (s *employeeService) Save(ctx context.Context,
	req *pb.EmployeeRequest) (*pb.EmployeeResponse, error) {
	return nil, nil
}

func (s *employeeService) SaveAll(stream pb.EmployeeService_SaveAllServer) error {
	return nil
}

func (s *employeeService) AddPhoto(stream pb.EmployeeService_AddPhotoServer) error {
	md, ok := metadata.FromContext(stream.Context())
	if ok {
		fmt.Printf("Receiving photo for badge number: %v\n", md["badgenumber"])
	}
	imgData := []byte{}
	for {
		data, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("Done, Final file size: ", len(imgData))
			return stream.SendAndClose(&pb.AddPhotoResponse{IsOk: true})
		}
		if err != nil {
			return err
		}
		fmt.Printf("Received %v bytes\n", len(data.Data))
		imgData = append(imgData, data.Data...)
	}
	return nil
}
