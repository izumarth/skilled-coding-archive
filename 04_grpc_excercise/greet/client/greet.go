package main

import (
	"context"
	"log"
	"time"

	pb "github.com/izumarth/skilled-coding-archive/04-grpc-excercise/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreet(c pb.GreetServiceClient, timeout time.Duration) {
	log.Println("doGreet was invoked")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req := &pb.GreetRequest{
		FirstName: "izumarth",
	}

	res, err := c.Greet(ctx, req)

	if err != nil {
		e, ok := status.FromError(err)

		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Println("Dead line exceede")
			} else {
				log.Println("error occured")
			}
		} else {
			log.Fatalf("Non Grpc Error %v\n", err)
		}
	}

	log.Printf("Greeting: %s", res.Result)
}
