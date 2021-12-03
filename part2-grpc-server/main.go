package main

import (
	"errors"
	"github.com/jxy90/GRPC/part2-grpc-server/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
)

const port = ":5001"

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err.Error())
	}
	creds, err := credentials.NewServerTLSFromFile("X509/cert.pem", "X509/key.pem")
	if err != nil {
		log.Fatalln(err.Error())
	}
	options := []grpc.ServerOption{grpc.Creds(creds)}
	server := grpc.NewServer(options...)
	pb.RegisterEmployeeServiceServer(server, new(employeeService))
	log.Println("gRPC Server start...", port)
	server.Serve(listen)
}

type employeeService struct {
}

func (e *employeeService) GetByNo(ctx context.Context, req *pb.GetByNoRequest) (*pb.EmployeeResponse, error) {
	for _, employee := range employees {
		if employee.No == req.No {
			return &pb.EmployeeResponse{Employee: &employee}, nil
		}
	}
	return nil, errors.New("employee not found")
}
func (e *employeeService) GetAll(*pb.GetAllRequest, pb.EmployeeService_GetAllServer) error {
	return nil
}
func (e *employeeService) AddPhoto(context.Context, *pb.AddPhotoRequest) (*pb.AddPhotoResponse, error) {
	return nil, nil
}

func (e *employeeService) Save(context.Context, *pb.EmployeeRequest) (*pb.EmployeeResponse, error) {
	return nil, nil
}

func (e *employeeService) SaveAll(pb.EmployeeService_SaveAllServer) error {
	return nil
}
