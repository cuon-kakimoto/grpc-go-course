package main

import (
	"fmt"
	"log"

	"github.com/cuon-kakimoto/grpc-go-course/sum/sumpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I'm client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("cound not connect :%w", err)
	}

	defer cc.Close()

	c := sumpb.NewSumServiceClient(cc)
	fmt.Printf("Created client: %f", c)

}
