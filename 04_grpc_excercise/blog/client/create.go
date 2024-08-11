package main

import (
	"context"
	"log"

	pb "github.com/izumarth/skilled-coding-archive/04-grpc-excercise/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("----- creatBlog was invoked ------")

	blog := &pb.Blog{
		AuthorId: "izumarth",
		Title:    "First Blog",
		Content:  "hogehoge",
	}

	res, err := c.CreateBlog(context.Background(), blog)
	if err != nil {
		log.Fatalf("Can't Create Blog %v\n", err)
	}

	log.Printf("----- creatBlog was finshed id: %s ------", res.Id)

	return res.Id
}
