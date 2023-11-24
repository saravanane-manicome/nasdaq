package service

import (
	"context"
	"github.com/saravanane-manicome/nasdaq/provider/protobuf/quote"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"math/rand"
)

type QuoteProviderService struct {
	Symbols map[string]float64
	quote.UnimplementedQuoteServiceServer
}

func (quoteProviderService *QuoteProviderService) GetQuote(_ context.Context, in *quote.QuoteRequest) (*quote.QuoteReply, error) {
	log.Printf("received pb request for symbol %s", in.GetSymbol())
	symbol := in.GetSymbol()
	q, exists := quoteProviderService.requestQuote(symbol)

	if !exists {
		return nil, status.Error(codes.NotFound, "symbol not registered")
	}
	return &quote.QuoteReply{Symbol: symbol, Quote: q}, nil
}

/*
*
Here is the function supposed to request an external data source
Because it would be too much for the training purpose of this project, this function simply
returns a random value
*/
func (quoteProviderService *QuoteProviderService) requestQuote(symbol string) (float64, bool) {
	if _, ok := quoteProviderService.Symbols[symbol]; ok {
		quoteProviderService.Symbols[symbol] += (rand.Float64() - 0.5) * 1000
		return quoteProviderService.Symbols[symbol], true
	}
	return 0, false
}
