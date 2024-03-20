package main

import (
	"fmt"
	"log"

	pb "proto/home/djax/Documents/MyFolder/protobuf"

	"github.com/golang/protobuf/proto"
)

func main() {

	address := &pb.Address{
		Street: "123 Main st",
		City:   "Washinton",
		State:  "USA",
		Zip:    "123456",
	}

	person := &pb.Person{
		Name:    "Nike",
		Age:     35,
		Address: address,
	}

	data, err := proto.Marshal(person)
	if err != nil {
		log.Println("Marshalling error", err)
	}

	p1 := &pb.Person{}
	err = proto.Unmarshal(data, p1)
	if err != nil {
		log.Println("Unmarshalling error", err)
	}

	fmt.Println("Name", p1.GetName())
	fmt.Println("Age", p1.GetAge())
	fmt.Println("Street", p1.GetAddress().GetStreet())
	fmt.Println("City", p1.GetAddress().GetCity())
	fmt.Println("State", p1.GetAddress().GetState())
	fmt.Println("Zip", p1.GetAddress().GetZip())
}
