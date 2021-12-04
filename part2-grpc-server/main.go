package main

import (
	"errors"
	"fmt"
	"github.com/jxy90/GRPC/part2-grpc-server/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"io"
	"log"
	"net"
	"time"
)

const port = ":5001"

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err.Error())
	}
	creds, err := credentials.NewServerTLSFromFile("X509/server.crt", "X509/server.key")
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
func (e *employeeService) GetAll(req *pb.GetAllRequest, stream pb.EmployeeService_GetAllServer) error {
	for _, employee := range employees {
		err := stream.Send(&pb.EmployeeResponse{Employee: &employee})
		if err != nil {
			log.Fatalln(err.Error())
		}
		time.Sleep(2 * time.Second)
	}
	return nil
}
func (e *employeeService) AddPhoto(stream pb.EmployeeService_AddPhotoServer) error {
	meta, ok := metadata.FromIncomingContext(stream.Context())
	if ok {
		fmt.Printf("employee : %v \n", meta["no"][0])
	}
	img := make([]byte, 0)
	for {
		data, err := stream.Recv()
		if err == io.EOF {
			fmt.Printf("photo size: %v \n", len(img))
			return stream.SendAndClose(&pb.AddPhotoResponse{IsOK: true})
		}
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		fmt.Printf("received: %v \n", len(data.Data))
		img = append(img, data.Data...)
	}
}

func (e *employeeService) Save(context.Context, *pb.EmployeeRequest) (*pb.EmployeeResponse, error) {
	return nil, nil
}

func (e *employeeService) SaveAll(stream pb.EmployeeService_SaveAllServer) error {
	for {
		empReq, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		employees = append(employees, *empReq.Employee)
		stream.Send(&pb.EmployeeResponse{Employee: empReq.Employee})
	}
	for _, employee := range employees {
		fmt.Println(employee)
	}
	return nil
}

func (e *employeeService) CreateToken(context.Context, *pb.TokenRequest) (*pb.TokenResponse, error) {
	return nil, nil
}
