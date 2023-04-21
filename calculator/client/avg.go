package main

import (
	"context"
	"log"
	"time"

	pb "github.com/alperergul/gRPC/calculator/proto"
)

func doAvg(c pb.CalculatorServiceClient){

	log.Println("doAvg was invoked")

	reqs := []*pb.AvgRequest{
		{
			Number: 2,
		},
		{
			Number: 4,
		},
		{
			Number: 6,
		},
	}

	stream, err := c.Avg(context.Background())

	if err != nil {
		log.Fatalf("Error while caling Avg function %v\n",err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response from Avg function: %v\n",err)
	}

	log.Printf("Avg: %v\n", res.Avg)

	
}