package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"github.com/cuon-kakimoto/grpc-go-course/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Calculator(ctx context.Context, req *calculatorpb.CalculatorRequest) (*calculatorpb.CalculatorResponse, error) {
	fmt.Printf("Calculator function was invoked with %w\n", req)
	a := req.GetCalculator().GetA()
	b := req.GetCalculator().GetB()

	result := a + b
	res := &calculatorpb.CalculatorResponse{
		Result: result,
	}
	return res, nil
}

func (*server) PrimeNumberDecomposition(req *calculatorpb.PrimeNumberRequest, stream calculatorpb.CalculatorService_PrimeNumberDecompositionServer) error {

	fmt.Printf("PrimeNumberDecomposition function was invoked with %v\n	", req)

	num := req.GetNumber()

	var k int64
	k = 2
	N := num
	for {
		if N <= 1 {
			break
		}

		if N%k == 0 {
			fmt.Printf("prime number div %v\n", k)
			// stream.Send(res)

			N = N / k
			res := &calculatorpb.PrimeNumberResponse{
				Result: k,
			}
			stream.Send(res)
			time.Sleep(1000 * time.Millisecond)

		} else {
			k = k + 1
		}
	}
	return nil

}
func (*server) ComputeAverage(stream calculatorpb.CalculatorService_ComputeAverageServer) error {
	fmt.Printf("LongGreet function was invoked with as streaming request")

	numberAdd := int64(0)
	count := int64(0)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			result := numberAdd / count
			// we have finished reading the client stream.
			return stream.SendAndClose(&calculatorpb.ComputeAverageResponse{
				Result: result,
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream %v", req)
		}
		number := req.GetNumber()
		numberAdd += number
		count++
	}

}

func main() {
	fmt.Println("Hello world")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listened: %v", err)
	}

	s := grpc.NewServer()

	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
