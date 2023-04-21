package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/alperergul/gRPC/greet/proto"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {

	log.Println("LongGreet function was invıked")

	res := ""

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream %v\n", err)
		}

		res += fmt.Sprintf("Hello %s!\n", req.FirstName)

	}
}
