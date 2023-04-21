package main

import (
	"io"
	"log"

	pb "github.com/alperergul/gRPC/greet/proto"
)


func (s *Server) GreetEveyone(stream pb.GreetService_GreetEveyoneServer) error{

	log.Println("GreetEveryone was invoked")

	for {

		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream %v\n", err)
		}

		res := "Hello" + req.FirstName + "!"
		
		err = stream.Send(&pb.GreetResponse{
			Result: res,
		})

		if err != nil {
			log.Fatalf("Error whie sending data to client: %v\n",err)
		}
	}

}