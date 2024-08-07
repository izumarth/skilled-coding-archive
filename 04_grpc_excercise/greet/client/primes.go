package main

import (
	"context"
	"io"
	"log"

	pb "github.com/izumarth/skilled-coding-archive/04-grpc-excercise/greet/proto"
)

func doPrimse(c pb.PrimesServiceClient) {
	log.Println("doPrimes was invoked")

	req := &pb.PrimesRequest{
		NumberOne: 120,
	}

	stream, err := c.Primes(context.Background(), req)

	if err != nil {
		log.Fatalf("Could not Primes%v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}
		log.Printf("Primes: %d", msg.Result)
	}

}
