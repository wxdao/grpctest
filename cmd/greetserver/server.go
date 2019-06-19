package main

import (
	"context"
	"grpctest/pkg/greet"
	"log"
	"time"

	"google.golang.org/grpc/metadata"
)

type greetServiceServer struct {
}

func (s *greetServiceServer) Echo(ctx context.Context, req *greet.EchoRequest) (*greet.EchoReply, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	log.Println("Echo is called. request:", req, "in-metadata:", md)

	return &greet.EchoReply{
		Message: req.Message,
	}, nil
}

func (s *greetServiceServer) Tick(req *greet.TickRequest, stream greet.GreetService_TickServer) error {
	md, _ := metadata.FromIncomingContext(stream.Context())
	log.Println("Tick is called. request:", req, "in-metadata:", md)

	for i := 0; int64(i) < req.Count; i++ {
		reply := greet.TickReply{
			Timestamp: time.Now().UnixNano(),
		}
		err := stream.Send(&reply)
		if err != nil {
			log.Println("Tick stream err:", err)
			break
		}
		log.Println("Tick stream sended. reply:", reply, "in-metadata:", md)
		time.Sleep(time.Duration(req.Interval))
	}

	log.Println("Tick ended. in-metadata:", md)
	return nil
}
