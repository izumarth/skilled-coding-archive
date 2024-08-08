package main

import (
	"log"
	"net"

	pb "github.com/izumarth/skilled-coding-archive/04-grpc-excercise/greet/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.GreetServiceServer
	pb.SumServiceServer
	pb.PrimesServiceServer
	pb.AverageServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listin on: %v", err)
	}

	log.Printf("Listning on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &Server{})
	pb.RegisterSumServiceServer(s, &Server{})
	pb.RegisterPrimesServiceServer(s, &Server{})
	pb.RegisterAverageServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
