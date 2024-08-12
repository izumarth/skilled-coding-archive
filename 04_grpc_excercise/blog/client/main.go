package main

import (
	"log"

	pb "github.com/izumarth/skilled-coding-archive/04-grpc-excercise/blog/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {

	opts := []grpc.DialOption{}
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.NewClient(addr, opts...) // grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()

	c := pb.NewBlogServiceClient(conn)
	// id := createBlog(c)
	// readBlog(c, id) // "66ba606cfc9329e0c3daa7b3"
	// readBlog(c, "aNonExistingId")
	// updateBlog(c, "66ba606cfc9329e0c3daa7b3")
	// listBlog(c)
	listBlog(c)
	deleteBlog(c, "66b928df96e95b06abc3fa9b")
	listBlog(c)
}
