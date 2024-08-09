package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/izumarth/skilled-coding-archive/04-grpc-excercise/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("doGreetEveryone was invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "izumarth"},
		{FirstName: "izumarth2"},
		{FirstName: "izumarth3"},
	}

	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Could not Long Great%v\n", err)
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
			log.Printf("Receved: %s", res.Result)
		}
		close(watic)
	}()

	<-watic
}
