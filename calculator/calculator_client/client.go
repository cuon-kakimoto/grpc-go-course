package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/cuon-kakimoto/grpc-go-course/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I'm client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("cound not connect :%w", err)
	}

	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)
	// fmt.Printf("Created client: %f", c)

	// doUnary(c)

	// doServerStreaming(c)

	// doClientStreaming(c)

	doBiDiStreaming(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {

	req := &calculatorpb.CalculatorRequest{
		Calculator: &calculatorpb.Calculator{
			A: 10,
			B: 3,
		},
	}
	res, err := c.Calculator(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling Greeet RPC: %v", err)
	}

	log.Printf("Response from Greet: %v", res.Result)

}

func doServerStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a Streming RPC...")

	req := &calculatorpb.PrimeNumberRequest{
		Number: 120,
	}

	resStream, err := c.PrimeNumberDecomposition(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling PrimeNumberDecomposition RPC: %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// we've reached the end of the stream
			break
		}

		if err != nil {
			log.Fatalf("error while reeading stream: %v", err)
		}
		log.Printf("Response from PrimeNumberDecomposition: %v", msg.GetResult())
	}

}

func doClientStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a Client Streming RPC...")

	requests := []*calculatorpb.ComputeAverageRequest{
		&calculatorpb.ComputeAverageRequest{
			Number: 1000,
		},
		&calculatorpb.ComputeAverageRequest{
			Number: 2,
		},
		&calculatorpb.ComputeAverageRequest{
			Number: 3,
		},
	}

	stream, err := c.ComputeAverage(context.Background())
	if err != nil {
		log.Fatalf("error while calling ComputeAverage RPC: %v", err)
	}

	// we iterate over our slice and send each message individualiy
	for _, req := range requests {
		fmt.Printf("Sending req: %v\n", req)
		stream.Send(req)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receving response ComputeAverage RPC: %v", err)
	}
	fmt.Printf("Sending req: %v", res)

}

func doBiDiStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a BiDi Streming RPC...")

	// we create a stream by invoking client
	stream, err := c.FindMaximum(context.Background())

	if err != nil {
		log.Fatalf("Error while creating stream:: %v", err)
		return
	}

	requests := []*calculatorpb.FindMaximumRequest{
		&calculatorpb.FindMaximumRequest{
			Number: 1,
		},
		&calculatorpb.FindMaximumRequest{
			Number: 5,
		},
		&calculatorpb.FindMaximumRequest{
			Number: 3,
		},
		&calculatorpb.FindMaximumRequest{
			Number: 6,
		},
		&calculatorpb.FindMaximumRequest{
			Number: 20,
		},
	}

	waitc := make(chan struct{})
	// we send a bunch of message the client(go ruotine)
	go func() {
		for _, req := range requests {
			fmt.Printf("Sending message: %v\n", req)
			stream.Send(req)
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()
	// we received a bunch of messages from the client(go routine)
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("error while receving: %v", err)
				break
			}
			fmt.Printf("Recevived: %v\n", res.GetResult())
		}
		close(waitc)
	}()

	// block until everything is done
	<-waitc

}
