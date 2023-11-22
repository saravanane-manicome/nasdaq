package main

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/saravanane-manicome/nasdaq/quote"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"net"
)

/*
*
Protobuf server implementation
*/
type server struct {
	pb.UnimplementedQuoteServiceServer
}

func (s *server) GetQuote(_ context.Context, in *pb.QuoteRequest) (*pb.QuoteReply, error) {
	log.Printf("received pb request for symbol %s", in.GetSymbol())
	symbol := in.GetSymbol()
	q := requestQuote(symbol)
	return &pb.QuoteReply{Symbol: symbol, Quote: q, Exists: true}, nil
}

/*
*
Here is the function supposed to request an external data source
Because it would be too much for the training purpose of this project, this function simply
returns a random value
*/
func requestQuote(symbol string) float64 {
	return rand.Float64() * 1000
}

func main() {
	port := flag.Int("port", 50051, "The server port")
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterQuoteServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
