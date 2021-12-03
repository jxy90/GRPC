package main

import (
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/jxy90/GRPC/part1-protobuf/src/complexpb"
	"github.com/jxy90/GRPC/part1-protobuf/src/enumpb"
	firstpb "github.com/jxy90/GRPC/part1-protobuf/src/firstpb"
	"io/ioutil"
	"log"
)

func main() {
	//first()
	//second()
	third()
}

func third() {
	dm := NewDepartmentMessage()
	fmt.Println(dm)
}

func NewDepartmentMessage() *complexpb.DepartmentMessage {
	dm := &complexpb.DepartmentMessage{
		Id:   5678,
		Name: "开发部",
		Employees: []*complexpb.EmployeeMessage{
			&complexpb.EmployeeMessage{
				Id:   11,
				Name: "Dave",
			},
			&complexpb.EmployeeMessage{
				Id:   11,
				Name: "Mike",
			},
		},
		ParentDepartment: &complexpb.DepartmentMessage{
			Id:   1,
			Name: "总部",
		},
		//ChildDepartment:  nil,
	}
	return dm
}

func second() {
	em := NewEnumMessage()
	fmt.Println(em)
	fmt.Println(enumpb.Gender_name[int32(em.Gender)])
}

func NewEnumMessage() *enumpb.EnumMessage {
	em := enumpb.EnumMessage{
		Id:     345,
		Gender: enumpb.Gender_FEMALE,
	}
	return &em
}

func first() {
	//fmt.Println(1)
	//pm := NewPersonMessage()
	//writeToFile("person.bin", pm)

	//pm2 := &firstpb.PersonMessage{}
	//readFromFile("person.bin", pm2)
	//fmt.Println(pm2)

	pm := NewPersonMessage()
	pmStr := toJson(pm)
	fmt.Println(pmStr)

	pm3 := &firstpb.PersonMessage{}
	fromJson(pmStr, pm3)
	fmt.Println(pm3)
}

func fromJson(jsonStr string, message proto.Message) error {
	err := jsonpb.UnmarshalString(jsonStr, message)
	if err != nil {
		log.Fatalln("fromJson Err", err.Error())
	}
	return err
}

func toJson(pb proto.Message) string {
	marshaler := jsonpb.Marshaler{Indent: " "}
	str, err := marshaler.MarshalToString(pb)
	if err != nil {
		return err.Error()
	}
	return str
}

func readFromFile(fileName string, pb proto.Message) error {
	databytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalln("read err", err.Error())
	}
	err = proto.UnmarshalMerge(databytes, pb)
	if err != nil {
		log.Fatalln("trans err", err.Error())
	}
	return nil
}

func writeToFile(fileName string, pb proto.Message) error {
	//序列化
	dataBytes, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("无法序列化到Bytes", err.Error())
		return err
	}
	if err := ioutil.WriteFile(fileName, dataBytes, 0644); err != nil {
		log.Fatalln("无法写入到文件", err.Error())
	}
	log.Println("写入到文件")
	return nil
}

func NewPersonMessage() *firstpb.PersonMessage {
	pm := firstpb.PersonMessage{
		Id:       1234,
		IsAdult:  true,
		Name:     "Dave",
		LuckNums: []int32{1, 2, 3, 4, 5, 6, 7, 8},
	}
	fmt.Println(pm)
	pm.Name = "Nick"
	fmt.Println(pm)
	fmt.Println(pm.GetId())
	return &pm
}
