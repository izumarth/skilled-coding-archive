package main

import (
	"context"
	"log"

	pb "github.com/izumarth/skilled-coding-archive/04-grpc-excercise/blog/proto"
)

func deleteBlog(c pb.BlogServiceClient, id string) {
	log.Println("----- deleteBlog was invoked ------")

	deleteId := &pb.BlogId{
		Id: id,
	}

	_, err := c.DeleteBlog(context.Background(), deleteId)
	if err != nil {
		log.Fatalf("Can't Delete Blog %v\n", err)
	}

	log.Println("----- deleteBlog was finshed")
}
