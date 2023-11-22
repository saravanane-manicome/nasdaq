package service

import (
	"context"
	pb "github.com/saravanane-manicome/nasdaq/quote"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func requestQuote(symbol string) func(pb.QuoteServiceClient, context.Context) (*Quote, error) {
	return func(client pb.QuoteServiceClient, context context.Context) (*Quote, error) {
		r, err := client.GetQuote(context, &pb.QuoteRequest{Symbol: symbol})
		if err != nil {
			log.Fatalf("could not request pb: %v", err)
			return nil, err
		}

		return &Quote{r.GetSymbol(), r.GetQuote(), r.GetExists()}, nil
	}
}

func setupClient(addr string) (pb.QuoteServiceClient, context.Context, func()) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	c := pb.NewQuoteServiceClient(conn)

	// Contact the server.
	ctx, cancel := context.WithCancel(context.Background())
	return c, ctx, func() {
		cancel()
		conn.Close()
	}
}

type QuoteService struct {
	Address string
}

func (service QuoteService) GetQuote(symbol string) (*Quote, error) {
	client, ctx, cancel := setupClient(service.Address)
	defer cancel()
	return requestQuote(symbol)(client, ctx)
}
