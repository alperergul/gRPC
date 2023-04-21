package main

import (
	"io"
	"log"

	pb "github.com/alperergul/gRPC/calculator/proto"
)

func (s *Server) Avg(stream pb.CalculatorService_AvgServer) error{


	log.Println("Avg function was invoked");


	var avg float32 = 0;
	var sum float32 = 0;
	count := 0;

	for {

		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.AvgResponse{
				Avg: float32(avg),
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream %v\n", err)
		}

		count += 1;
		sum += float32(req.Number)
		avg = sum / float32(count)
	}


}