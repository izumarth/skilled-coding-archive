package main

import (
	"context"
	"log"

	pb "github.com/izumarth/skilled-coding-archive/04-grpc-excercise/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "izumarth",
	})

	if err != nil {
		log.Fatalf("Could not Greet %v\n", err)
	}

	log.Printf("Greeting: %s", res.Result)
}

func doSum(c pb.SumServiceClient) {
	log.Println("dSum was invoked")
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		NumberOne: 8,
		NumberTwo: 5,
	})

	if err != nil {
		log.Fatalf("Could not Sum %v\n", err)
	}

	log.Printf("Sum: %d", res.Result)
}
