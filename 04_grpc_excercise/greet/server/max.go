package main

import (
	"io"
	"log"

	pb "github.com/izumarth/skilled-coding-archive/04-grpc-excercise/greet/proto"
)

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Printf("MaxService_MaxServer function was invoked")

	max := int64(-1)
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream %v\n", err)
		}

		log.Printf("Receiving: %v\n", req)

		received_number := req.Number

		if received_number > max {
			max = received_number

			err = stream.Send(&pb.MaxResponse{
				Result: max,
			})

			if err != nil {
				log.Fatalf("Error while sending client stream %v\n", err)
			}
		}
	}
}
