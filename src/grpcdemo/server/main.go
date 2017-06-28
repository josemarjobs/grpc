package main

import (
	"fmt"
	"grpcdemo/pb"
	"log"
	"net"

	"google.golang.org/grpc/credentials"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
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
	return &pb.EmployeeResponse{
		Employee: &employees[0],
	}, nil
}

func (s *employeeService) GetAll(req *pb.GetAllRequest,
	stream pb.EmployeeService_GetAllServer) error {
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
	return nil
}
