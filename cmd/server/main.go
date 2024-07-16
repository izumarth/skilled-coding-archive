package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

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

func (s *myServer) HelloServerStream(
	req *examplepb.HelloRequest,
	stream examplepb.GreetingService_HelloServerStreamServer,
) error {
	resCount := 5
	for i := 0; i < resCount; i++ {
		if err := stream.Send(
			&examplepb.HelloResponse{
				Message: fmt.Sprintf("[%d] Hello, %s!", i, req.GetName()),
			}); err != nil {
			return err
		}
		time.Sleep(time.Second * 1)
	}
	return nil
}

func (s *myServer) HelloClientStream(
	stream examplepb.GreetingService_HelloClientStreamServer,
) error {
	nameList := make([]string, 0)
	for {
		req, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			message := fmt.Sprintf("Hello %v", nameList)
			return stream.SendAndClose(
				&examplepb.HelloResponse{
					Message: message,
				},
			)
		}

		if err != nil {
			return err
		}
		nameList = append(nameList, req.GetName())
	}
}

func (s *myServer) HelloBiStream(
	stream examplepb.GreetingService_HelloBiStreamServer,
) error {
	for {
		req, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			return nil
		}
		if err != nil {
			return err
		}

		message := fmt.Sprintf("Hello, %v!", req.GetName())
		if err := stream.Send(
			&examplepb.HelloResponse{
				Message: message,
			},
		); err != nil {
			return nil
		}
	}
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
