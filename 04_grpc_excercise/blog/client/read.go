package main

import (
	"context"
	"log"

	pb "github.com/izumarth/skilled-coding-archive/04-grpc-excercise/blog/proto"
)

func readBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	log.Println("----- readBlog was invoked ------")

	req := &pb.BlogId{
		Id: id,
	}

	res, err := c.ReadBlog(context.Background(), req)
	if err != nil {
		log.Fatalf("Can't Read Blog %v\n", err)
	}

	log.Printf("Blog was read: %v\n", res)

	return res
}
