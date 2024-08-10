package main

import (
	"context"
	"log"
	"time"

	pb "github.com/izumarth/skilled-coding-archive/04-grpc-excercise/greet/proto"
	"google.golang.org/grpc/status"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doLongGreet was invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "izumarth"},
		{FirstName: "izumarth2"},
		{FirstName: "izumarth3"},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		res, ok := status.FromError(err)
		if ok {
			log.Printf("Error Message From Server: %v", res.Message())
		} else {
			log.Fatalf("Could not Long Great%v\n", err)
		}
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

	log.Printf("LongGreet %s\n", res.Result)
}
