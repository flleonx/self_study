package main

import (
	"fmt"
	pb "grpc-demo/proto"
	"io"
	"log"
)

func (s *helloServer) SayHelloBidirectionalStreaming(stream pb.GreetService_SayHelloBidirectionalStreamingServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		log.Printf("Got request with name: %s", req.Name)
		res := &pb.HelloResponse{
			Message: fmt.Sprintf("Hello %s", req.Name),
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}
}
