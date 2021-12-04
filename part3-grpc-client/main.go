package main

import (
	"context"
	"fmt"
	"github.com/jxy90/GRPC/part3-grpc-client/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
	"io"
	"log"
	"os"
	"time"
)

const port = ":5001"

func main() {
	creds, err := credentials.NewClientTLSFromFile("X509/ca.crt", "www.lixueduan.com")
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
	//getByNo(client)
	getAll(client)
	addPhoto(client)
	//saveAll(client)
}

//双向
func saveAll(client pb.EmployeeServiceClient) {
	employees := []pb.Employee{
		{
			Id:        10,
			No:        1994,
			FirstName: "Chandler",
			LastName:  "Bing",
			MouthSalary: &pb.MouthSalary{
				Basic: 5000,
				Bonus: 125.5,
			},
			Status: pb.EmployeeStatus_NORMAL,
			LastModified: &timestamppb.Timestamp{
				Seconds: time.Now().Unix(),
			},
		},
		{
			Id:        11,
			No:        1995,
			FirstName: "Chang",
			LastName:  "Jin",
			MouthSalary: &pb.MouthSalary{
				Basic: 10000,
				Bonus: 625.5,
			},
			Status: pb.EmployeeStatus_NORMAL,
			LastModified: &timestamppb.Timestamp{
				Seconds: time.Now().Unix(),
			},
		},
	}
	stream, err := client.SaveAll(context.Background())
	if err != nil {
		log.Fatalln(err.Error())
	}
	finishChan := make(chan struct{})
	go func() {
		for {
			fmt.Println("------")
			res, err := stream.Recv()
			if err == io.EOF {
				fmt.Println("finish")
				finishChan <- struct{}{}
				break
			}
			if err != nil {
				log.Fatalln(err.Error())
			}
			fmt.Println(res.Employee)
		}
	}()
	for _, employee := range employees {
		stream.Send(&pb.EmployeeRequest{Employee: &employee})
	}
	stream.CloseSend()
	<-finishChan
}

//Client Streaming
func addPhoto(client pb.EmployeeServiceClient) {
	fmt.Println("into addPhoto")
	imgFile, err := os.Open("1.jpg")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer imgFile.Close()
	meta := metadata.New(map[string]string{"no": "1994"})
	context := context.Background()
	context = metadata.NewOutgoingContext(context, meta)

	stream, err := client.AddPhoto(context)
	if err != nil {
		log.Fatalln(err.Error())
	}
	for {
		chunk := make([]byte, 128*1024)
		chunkSize, err := imgFile.Read(chunk)
		if err == io.EOF {
			log.Println("file read over")
			break
		}
		if err != nil {
			log.Fatalln(err.Error())
		}
		fmt.Printf("send chunk: %v \n", chunkSize)
		if chunkSize < len(chunk) {
			fmt.Println("file send over")
			chunk = chunk[:chunkSize]
		}
		stream.Send(&pb.AddPhotoRequest{Data: chunk})
		time.Sleep(time.Millisecond * 500)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(res)
}

//Server Streaming
func getAll(client pb.EmployeeServiceClient) {
	stream, err := client.GetAll(context.Background(), &pb.GetAllRequest{})
	if err != nil {
		log.Fatalln(err.Error())
	}
	for {
		rev, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				//log.Println(err.Error())
				return
			}
			log.Fatalln(err.Error())
		}
		fmt.Println(rev)
	}
}

//一元消息
func getByNo(client pb.EmployeeServiceClient) {
	res, err := client.GetByNo(context.Background(), &pb.GetByNoRequest{No: 1994})
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(res.Employee)
}
