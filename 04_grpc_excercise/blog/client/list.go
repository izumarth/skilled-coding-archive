package main

import (
	"context"
	"io"
	"log"

	pb "github.com/izumarth/skilled-coding-archive/04-grpc-excercise/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func listBlog(c pb.BlogServiceClient) {
	log.Println("----- listBlog was invoked ------")

	stream, err := c.List(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf("Error while calling listBlog %v\n", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error ocuured %v\n", err)
		}

		log.Println(res)
	}
}
