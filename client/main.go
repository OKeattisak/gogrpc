package main

import (
	"client/services"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	creds := insecure.NewCredentials()
	cc, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	calculatorClient := services.NewCalculatorClient(cc)
	calculatorService := services.NewCalculatorService(calculatorClient)

	// err = calculatorService.Hello("Keattisak")
	// err = calculatorService.Fibonacci(6)
	// err = calculatorService.Average(1, 2, 3, 4, 5)
	err = calculatorService.Sum(1, 2, 3, 4, 5)

	if err != nil {
		log.Fatal(err)
	}
}
