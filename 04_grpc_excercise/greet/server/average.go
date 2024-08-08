package main

import (
	"io"
	"log"

	pb "github.com/izumarth/skilled-coding-archive/04-grpc-excercise/greet/proto"
)

func (s *Server) Average(stream pb.AverageService_AverageServer) error {
	log.Println("Average function was invoked")

	res := []int64{}

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.AverageResponse{
				Result: average(res),
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream %v\n", err)
		}

		res = append(res, req.Number)
	}
}

func average(arr []int64) float32 {
	sum := float32(0)
	for _, num := range arr {
		sum += float32(num)
	}
	return sum / float32(len(arr))
}
