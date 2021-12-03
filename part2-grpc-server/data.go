package main

import (
	"github.com/jxy90/GRPC/part2-grpc-server/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

var employees = []pb.Employee{
	{
		Id:        1,
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
		Id:        2,
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
