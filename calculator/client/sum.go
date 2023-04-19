package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/alperergul/gRPC/calculator/proto"
)

func doCalculate(c pb.CalculatorServiceClient) {

	log.Println("doCalculate is invoked.")

	res, err := c.Calculate(context.Background(), &pb.CalculatorRequest{
		FirstNumber:  10,
		SecondNumber: 3,
	})

	if err != nil {
		log.Fatalf("Could not calculate: %v\n", err)
	}

	fmt.Printf("10 + 3 = %v", res.Result)
}
