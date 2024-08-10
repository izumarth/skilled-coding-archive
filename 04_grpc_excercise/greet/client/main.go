package main

import (
	"log"
	"time"

	pb "github.com/izumarth/skilled-coding-archive/04-grpc-excercise/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)
	doGreet(c, 1*time.Second)
	//doGreetManyTimes(c)
	//doLongGreet(c)

	//c_sum := pb.NewSumServiceClient(conn)
	//doSum(c_sum)

	//c := pb.NewPrimesServiceClient(conn)
	//doPrimse(c)

	// c := pb.NewAverageServiceClient(conn)
	// doAverage(c)

	//c := pb.NewCalculatorServiceClient(conn)
	//doMax(c)

}
