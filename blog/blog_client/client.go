package main

import (
	"context"
	"fmt"
	"log"

	"github.com/cuon-kakimoto/grpc-go-course/blog/blogpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Blog Client")

	opts := grpc.WithInsecure()

	cc, err := grpc.Dial("localhost:50051", opts)

	if err != nil {
		log.Fatalf("counld not connect: %v", err)
	}

	defer cc.Close()

	c := blogpb.NewBlogServiceClient(cc)

	// create blog
	fmt.Println("Creating the blog")
	blog := &blogpb.Blog{
		AuthorId: "shinji",
		Title:    "My first Blog",
		Content:  "Contents",
	}
	createBlogRes, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: blog})
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}
	fmt.Printf("Blog has been created: %v", createBlogRes)
	blogID := createBlogRes.GetBlog().GetId()

	// read blog
	fmt.Println("Reading the blog")

	// for NOT FOUND
	_, err2 := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{BlogId: "11113jkj"})
	if err2 != nil {
		fmt.Printf("Error happend while reading: %v \n", err2)
	}

	readBlogReq := &blogpb.ReadBlogRequest{BlogId: blogID}
	readBlogRes, readBlogErr := c.ReadBlog(context.Background(), readBlogReq)
	if readBlogErr != nil {
		fmt.Printf("Error happend while reading: %v \n", err2)
	}

	fmt.Printf("Blog was read: %v", readBlogRes)

	// update Blog

	newBlog := &blogpb.Blog{
		Id:       blogID,
		AuthorId: "Changed Author",
		Title:    "My first Blog (edited)",
		Content:  "Content of ther first blog, with some additions",
	}
	updateRes, updateErr := c.UpdateBlog(context.Background(), &blogpb.UpdateBlogRequest{Blog: newBlog})
	if updateErr != nil {
		fmt.Printf("Error happend while updating: %v \n", err2)
	}

	fmt.Printf("Blog was read: %v", updateRes)

}
