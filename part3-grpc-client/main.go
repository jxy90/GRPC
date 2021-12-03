package main

import (
	"context"
	"fmt"
	"github.com/jxy90/GRPC/part3-grpc-client/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

const port = ":5001"

func main() {
	creds, err := credentials.NewClientTLSFromFile("cert.pem", "")
	if err != nil {
		log.Fatalln(err.Error())
	}
	options := []grpc.DialOption{grpc.WithTransportCredentials(creds)}
	conn, err := grpc.Dial("localhost"+port, options...)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer conn.Close()
	client := pb.NewEmployeeServiceClient(conn)
	GetByNo(client)
}

func GetByNo(client pb.EmployeeServiceClient) {
	res, err := client.GetByNo(context.Background(), &pb.GetByNoRequest{No: 1999})
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(res.Employee)
}
