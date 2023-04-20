package main

import (
	"context"
	"io"
	"log"

	pb "github.com/alperergul/gRPC/calculator/proto"
)

func doPrimes(c pb.CalculatorServiceClient) {

	log.Printf("doPrimes was invoked\n")

	req := &pb.PrimeRequest{
		Number: 12390392840,
	}

	stream, err := c.Primes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling Primes: %v\n", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading stream: %v\n", err)
		}

		log.Printf("Primes: %v\n", res.Result)
	}

}
