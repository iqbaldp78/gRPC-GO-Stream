package main

import (
	"context"
	"fmt"
	"log"
	"strconv"

	pb "grpc-go/modules/pb"

	"google.golang.org/grpc"
)

var (
	ServerAddr string = "localhost:10000"
)

func main() {
	startClientPush()
}

//startClient used for test client conn
func startClientPush() error {
	log.Println("starting push .....")
	var (
		num string
	)
	conn, err := grpc.Dial(ServerAddr, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewAvgAlarmServiceClient(conn)

	fmt.Println("Enter Number to push: ")
	fmt.Scanln(&num)

	numInt, err := strconv.Atoi(num)
	if err != nil {
		log.Fatal("please enter number only")
	}
	data := &pb.PushReq{Num: float32(numInt)}

	_, err = client.Push(context.Background(), data)
	if err != nil {
		log.Fatalf("%v.Push(_) = _, %v: ", client, err)
	}
	return nil
}
