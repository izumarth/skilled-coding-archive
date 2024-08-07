package main

import (
	"context"
	"io"
	"log"

	pb "github.com/izumarth/skilled-coding-archive/04-grpc-excercise/greet/proto"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Println("doGreetManyTimes was invoked")

	req := &pb.GreetRequest{
		FirstName: "izumarth",
	}

	stream, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("Could not Greet Many Times%v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}
		log.Printf("Greeting Many Times: %s", msg.Result)

	}

}
