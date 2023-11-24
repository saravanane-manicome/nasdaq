package main

import (
	"flag"
	"fmt"
	"github.com/saravanane-manicome/nasdaq/provider/service"
	pb "github.com/saravanane-manicome/nasdaq/quote"
	"google.golang.org/grpc"
	"log"
	"net"
)

/*
*
Protobuf QuoteService implementation
*/

func main() {
	port := flag.Int("port", 50051, "The server port")
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	symbols := map[string]float64{
		"AMD":      1000,
		"INTEL":    1000,
		"QUALCOMM": 1000,
	}
	s := grpc.NewServer()
	pb.RegisterQuoteServiceServer(s, &service.QuoteProviderService{Symbols: symbols})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
