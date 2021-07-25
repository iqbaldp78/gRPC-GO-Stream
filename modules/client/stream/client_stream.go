package main

import (
	"context"
	"fmt"
	pb "grpc-go/modules/pb"
	"io"
	"log"
	"strconv"

	"google.golang.org/grpc"
)

var (
	ServerAddr string = "localhost:10000"
)

func main() {
	startClientStream()
}

func startClientStream() {
	log.Println("starting client stream ....")
	var (
		num string
	)

	// Create our context
	ctx := context.Background()

	// Setup connection
	conn, err := grpc.DialContext(ctx, ServerAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	// Close connection when we are done
	defer conn.Close()

	// Use the generated NewAvgAlarmServiceClient method and pass our Connection
	client := pb.NewAvgAlarmServiceClient(conn)

	fmt.Println("Enter Number of treshold: ")
	fmt.Scanln(&num)

	numInt, err := strconv.Atoi(num)
	if err != nil {
		log.Fatal("please enter number only")
	}
	// Call rquest to receive the Stream of data
	req := &pb.GetAvgAlarmStreamReq{Threshold: float32(numInt)}

	// call function, this will return a stream of data
	stream, err := client.GetAvgAlarmStream(ctx, req)
	if err != nil {
		panic(err)
	}

	done := make(chan bool)

	go func() {
		for {
			// Start receiving streaming messages
			res, err := stream.Recv()
			if err == io.EOF {
				// break
				done <- true //means stream is finished
				return
			}
			if err != nil {
				log.Fatalf("error when receiving server response stream: %v", err)
			}
			log.Printf("AVG value is %v greather than treshold:", res.Avg)
			// break
		}
	}()
	done <- true //means stream is finished
	return
}
