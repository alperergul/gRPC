package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/alperergul/gRPC/greet/proto"
)

func doGreetEverone(c pb.GreetServiceClient){

	log.Println("doGreetEveyone was invoked")

	stream, err := c.GreetEveyone(context.Background())

	if err != nil {
		log.Fatalf("Error while creating stream: %v\n",err)
	}

	reqs := []*pb.GreetRequest{
		{
			FirstName: "Alper",
		},
		{
			FirstName: "Feyza",
		},
		{
			FirstName: "Ezgi",
		},
	}

	waitc := make(chan struct{})

	go func ()  {
		for _, req := range reqs {
			log.Printf("Send request: %v\n",req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func ()  {
		
		for {
			res, err := stream.Recv()

			if err == io.EOF{
				break
			}

			if err != nil {
				log.Printf("Error while reciving: %v\n",err)
			}

			log.Printf("Received: %v\n", res.Result)
		}
		close(waitc)
	}()

	<-waitc

} 