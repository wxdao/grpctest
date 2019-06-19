package main

import (
	"context"
	"flag"
	"fmt"
	"grpctest/pkg/greet"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
)

var greetURL = flag.String("greet_url", "localhost:8000", "Greet service endpoint")

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Usage:", os.Args[0], "[-greet_url=localhost:8000]", "<action>")
		fmt.Println()
		fmt.Println("Available actions:")
		fmt.Println("loopecho")
		return
	}

	switch args[0] {
	case "loopecho":
		loopEcho()
	case "tick":
		tick()
	default:
		log.Println("unknown action")
	}
}

func loopEcho() {
	greetConn, err := grpc.Dial(*greetURL, grpc.WithInsecure())
	if err != nil {
		log.Fatal("failed to dial greet service:", err)
	}
	greetClient := greet.NewGreetServiceClient(greetConn)
	log.Println("dailed greet service")

	for {
		log.Println("calling echo")
		reply, err := greetClient.Echo(context.Background(), &greet.EchoRequest{
			Message: time.Now().String(),
		})
		if err != nil {
			log.Println("got err:", err)
		} else {
			log.Println("got reply:", reply)
		}
		time.Sleep(time.Millisecond * 1000)
	}
}

func tick() {
	greetConn, err := grpc.Dial(*greetURL, grpc.WithInsecure())
	if err != nil {
		log.Fatal("failed to dial greet service:", err)
	}
	greetClient := greet.NewGreetServiceClient(greetConn)
	log.Println("dailed greet service")

	for {
		log.Println("calling tick")
		stream, err := greetClient.Tick(context.Background(), &greet.TickRequest{
			Count:    100000000,
			Interval: int64(time.Second),
		})
		if err != nil {
			log.Println("call got err:", err)
			log.Println("will try again in 1s")
			time.Sleep(time.Second)
			continue
		}
		for {
			reply, err := stream.Recv()
			if err != nil {
				log.Println("stream recv got err:", err)
				break
			} else {
				log.Println("stream recv got reply:", reply)
			}
		}
	}
}
