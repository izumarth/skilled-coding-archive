package main

import (
	"log"

	pb "github.com/izumarth/skilled-coding-archive/04-grpc-excercise/greet/proto"
)

func (s *Server) Primes(in *pb.PrimesRequest, stream pb.CalculatorService_PrimesServer) error {
	log.Printf("Primes function was invoked with: %v\n", in)

	factors := factorize(in.NumberOne)

	for _, factor := range factors {
		stream.Send(&pb.PrimesResponse{Result: factor})
	}

	return nil
}

func factorize(n int64) []int64 {
	var factors []int64
	k := int64(2)

	for n > 1 {
		if n%k == 0 {
			factors = append(factors, k)
			n = n / k
		} else {
			k = k + 1
		}

	}

	return factors
}
