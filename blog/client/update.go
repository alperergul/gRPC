package main

import (
	"context"
	"log"

	pb "github.com/alperergul/gRPC/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {

	log.Println("---updateBlog was invoked---")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Not Alper",
		Title:    "A new title",
		Content:  "Content of the first blog, with some awesome additonal",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)

	if err != nil {
		log.Fatalf("Error happened while updating: %v\n", err)
	}

	log.Println("Blog was updated")

}
