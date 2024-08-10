package main

import (
	"log"

	pb "github.com/izumarth/skilled-coding-archive/04-grpc-excercise/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {

	tls := true
	opts := []grpc.DialOption{}

	if tls {
		certFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")

		if err != nil {
			log.Fatalf("Error while loading CA trust %v\n", err)
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	conn, err := grpc.NewClient(addr, opts...) // grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()

	//c := pb.NewGreetServiceClient(conn)
	//doGreet(c, 5*time.Second)
	//doGreetManyTimes(c)
	//doLongGreet(c)

	//c_sum := pb.NewSumServiceClient(conn)
	//doSum(c_sum)

	//c := pb.NewPrimesServiceClient(conn)
	//doPrimse(c)

	// c := pb.NewAverageServiceClient(conn)
	// doAverage(c)

	c := pb.NewCalculatorServiceClient(conn)
	doMax(c)

}
