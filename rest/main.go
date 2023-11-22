package main

import (
	"flag"
	"github.com/saravanane-manicome/nasdaq/rest/controller"
	"github.com/saravanane-manicome/nasdaq/rest/service"
)

func main() {
	address := flag.String("address", "localhost:50051", "the address to connect to")
	flag.Parse()

	quoteController := controller.QuoteController{
		QuoteService: service.QuoteService{
			Address: *address,
		},
	}

	quoteController.Serve()
}
