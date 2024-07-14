package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	examplepb "izumarth_grpc/pkg/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type myServer struct {
	examplepb.UnimplementedGreetingServiceServer
}

func (s *myServer) Hello(
	ctx context.Context,
	req *examplepb.HelloRequest,
) (*examplepb.HelloResponse, error) {
	return &examplepb.HelloResponse{
		Message: fmt.Sprintf("Hello, %s!", req.GetName()),
	}, nil
}

func NewMyServer() *myServer {
	return &myServer{}
}

func main() {
	port := 8080
	listner, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	examplepb.RegisterGreetingServiceServer(s, NewMyServer())

	reflection.Register(s)

	go func() {
		log.Printf("start gRPC server port: %v", port)
		s.Serve(listner)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()
}
