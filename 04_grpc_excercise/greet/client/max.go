package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/izumarth/skilled-coding-archive/04-grpc-excercise/greet/proto"
)

func doMax(c pb.MaxServiceClient) {
	log.Println("doMax was invoked")

	reqs := []*pb.MaxRequest{
		{Number: 1},
		{Number: 5},
		{Number: 3},
		{Number: 6},
		{Number: 2},
		{Number: 20},
	}

	stream, err := c.Max(context.Background())
	if err != nil {
		log.Fatalf("Could not Max%v\n", err)
	}

	watic := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Sending req: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Error while reading the stream: %v\n", err)
				break
			}
			log.Printf("Receved: %d", res.Result)
		}
		close(watic)
	}()

	<-watic
}
