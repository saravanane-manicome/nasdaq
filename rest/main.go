package main

import (
	"flag"
	"github.com/saravanane-manicome/nasdaq/rest/controller"
	"github.com/saravanane-manicome/nasdaq/rest/service"
)

func main() {
	providerAddress := flag.String("provider", "localhost:50051", "the address of the provider")
	flag.Parse()

	quoteController := controller.QuoteController{
		QuoteService: service.QuoteService{
			ProviderAddress: *providerAddress,
		},
	}

	quoteController.Serve()
}
