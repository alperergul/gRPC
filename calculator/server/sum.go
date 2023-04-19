package main

import (
	"context"
	"log"

	pb "github.com/alperergul/gRPC/calculator/proto"
)

func (s *Server) Calculate(ctx context.Context, in *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	log.Printf("Calculate function was invoked with %v\n", in)

	return &pb.CalculatorResponse{
		Result: in.FirstNumber + in.SecondNumber,
	}, nil
}
