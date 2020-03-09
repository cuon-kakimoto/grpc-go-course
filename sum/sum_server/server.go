package main

import (
	"fmt"
	"log"
	"net"

	"github.com/cuon-kakimoto/grpc-go-course/sum/sumpb"
	"google.golang.org/grpc"
)

type server struct{}

func main() {
	fmt.Println("Hello world")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listened: %v", err)
	}

	s := grpc.NewServer()

	sumpb.RegisterSumServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
