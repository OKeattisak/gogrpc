package services

import (
	context "context"
	"fmt"
	"io"
	"time"

	grpc "google.golang.org/grpc"
)

type calculatorServer struct {
}

func NewCalculatorServer() CalculatorServer {
	return calculatorServer{}
}

func (calculatorServer) mustEmbedUnimplementedCalculatorServer() {}

func (calculatorServer) Hello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	result := fmt.Sprintf("Hello %v at %v", req.Name, req.CreatedDate.AsTime().Local())
	res := HelloResponse{
		Result: result,
	}
	return &res, nil
}

func (calculatorServer) Fibonacci(req *FibonacciRequest, stream grpc.ServerStreamingServer[FibonacciResponse]) error {
	for n := uint32(0); n <= req.N; n++ {
		result := fib(n)
		res := FibonacciResponse{
			Result: result,
		}

		stream.Send(&res)
		time.Sleep(time.Second)
	}
	return nil
}

func fib(n uint32) uint32 {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return fib(n-1) + fib(n-2)
	}
}

func (calculatorServer) Average(stream grpc.ClientStreamingServer[AverageRequest, AverageResponse]) error {
	sum := 0.0
	count := 0.0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		sum += req.Number
		count++
	}

	res := AverageResponse{
		Result: sum / count,
	}

	return stream.SendAndClose(&res)
}
