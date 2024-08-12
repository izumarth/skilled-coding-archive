package main

import (
	"context"
	"log"

	pb "github.com/izumarth/skilled-coding-archive/04-grpc-excercise/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("----- updateBlog was invoked ------")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "izumarth",
		Title:    "Updated Blog",
		Content:  "hogehoge updated",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)
	if err != nil {
		log.Fatalf("Can't Update Blog %v\n", err)
	}

	log.Println("----- updateBlog was finshed")

	return
}
