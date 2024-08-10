package main

import (
	"context"
	"log"
	"time"

	pb "github.com/izumarth/skilled-coding-archive/04-grpc-excercise/greet/proto"
)

func doAverage(c pb.CalculatorServiceClient) {
	log.Println("doAverage was invoked")

	reqs := []*pb.AvrageRequest{
		{Number: 1},
		{Number: 2},
		{Number: 3},
		{Number: 4},
	}

	stream, err := c.Average(context.Background())
	if err != nil {
		log.Fatalf("Could not Average%v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receive response %v\n", err)
	}

	log.Printf("Average %d\n", res.Result)
}
