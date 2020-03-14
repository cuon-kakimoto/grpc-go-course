package main

import (
	"context"
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
	// fmt.Printf("Created client: %f", c)

	doUnary(c)
}

func doUnary(c sumpb.SumServiceClient) {

	req := &sumpb.SumRequest{
		Sum: &sumpb.Sum{
			A: 10,
			B: 3,
		},
	}
	res, err := c.Sum(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling Greeet RPC: %v", err)
	}

	log.Printf("Response from Greet: %v", res.Result)

}
