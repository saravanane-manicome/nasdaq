package main

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/saravanane-manicome/nasdaq/quote"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"math/rand"
	"net"
)

/*
*
Protobuf QuoteService implementation
*/
type QuoteService struct {
	symbols map[string]float64
	pb.UnimplementedQuoteServiceServer
}

func (quoteService *QuoteService) GetQuote(_ context.Context, in *pb.QuoteRequest) (*pb.QuoteReply, error) {
	log.Printf("received pb request for symbol %s", in.GetSymbol())
	symbol := in.GetSymbol()
	q, exists := quoteService.requestQuote(symbol)

	if !exists {
		return nil, status.Error(codes.NotFound, "symbol not registered")
	}
	return &pb.QuoteReply{Symbol: symbol, Quote: q}, nil
}

/*
*
Here is the function supposed to request an external data source
Because it would be too much for the training purpose of this project, this function simply
returns a random value
*/
func (quoteService *QuoteService) requestQuote(symbol string) (float64, bool) {
	if _, ok := quoteService.symbols[symbol]; ok {
		quoteService.symbols[symbol] += (rand.Float64() - 0.5) * 1000
		return quoteService.symbols[symbol], true
	}
	return 0, false
}

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
	pb.RegisterQuoteServiceServer(s, &QuoteService{symbols: symbols})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
