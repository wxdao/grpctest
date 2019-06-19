package main

import (
	"flag"
	"grpctest/pkg/greet"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"

	"google.golang.org/grpc"
)

var greetAddr = flag.String("greet_addr", "0.0.0.0:8000", "Listen address of greet service")
var pprofAddr = flag.String("pprof_addr", "0.0.0.0:6000", "Listen address of http pprof")

func main() {
	flag.Parse()

	errc := make(chan error)

	go func() {
		log.Println("started http pprof on", *pprofAddr)
		errc <- http.ListenAndServe(*pprofAddr, nil)
	}()

	greetLis, err := net.Listen("tcp", *greetAddr)
	if err != nil {
		log.Fatalln(err)
	}
	greetGRPCServer := grpc.NewServer()
	greet.RegisterGreetServiceServer(greetGRPCServer, &greetServiceServer{})
	go func() {
		log.Println("started greet service server on", *greetAddr)
		errc <- greetGRPCServer.Serve(greetLis)
	}()

	log.Fatal(<-errc)
}
