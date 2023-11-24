package service

import (
	"context"
	"github.com/saravanane-manicome/nasdaq/rest/protobuf/quote"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"log"
)

func requestQuote(symbol string) func(quote.QuoteServiceClient, context.Context) (*Quote, error) {
	return func(client quote.QuoteServiceClient, context context.Context) (*Quote, error) {
		r, err := client.GetQuote(context, &quote.QuoteRequest{Symbol: symbol})
		if err != nil {
			return handleRequestError(err)
		}

		return &Quote{r.GetSymbol(), r.GetQuote()}, nil
	}
}

func handleRequestError(err error) (*Quote, error) {
	responseStatus, ok := status.FromError(err)
	if ok && responseStatus.Code() == codes.NotFound {
		return nil, nil
	}
	log.Printf("could not request pb: %v\n", err)
	return nil, err
}

func setupClient(addr string) (quote.QuoteServiceClient, context.Context, func()) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("did not connect: %v\n", err)
	}

	c := quote.NewQuoteServiceClient(conn)

	// Contact the server.
	ctx, cancel := context.WithCancel(context.Background())
	return c, ctx, func() {
		cancel()
		conn.Close()
	}
}

type QuoteService struct {
	ProviderAddress string
}

func (service QuoteService) GetQuote(symbol string) (*Quote, error) {
	client, ctx, cancel := setupClient(service.ProviderAddress)
	defer cancel()
	return requestQuote(symbol)(client, ctx)
}
