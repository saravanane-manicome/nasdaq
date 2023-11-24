package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type NasdaqQuote struct {
	Symbol string  `json:"symbol"`
	Quote  float64 `json:"quote"`
}

func requestQuote(symbol string) (*NasdaqQuote, error) {
	response, err := http.Get(fmt.Sprintf("http://localhost:8080/quote?symbol=%s", symbol))

	if err != nil {
		return nil, err
	}

	switch response.StatusCode {
	case http.StatusOK:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		var quote NasdaqQuote
		err = json.Unmarshal(body, &quote)
		if err != nil {
			return nil, err
		}
		return &quote, nil
	case http.StatusNotFound:
		return nil, errors.New("symbol not found")
	case http.StatusBadRequest:
		return nil, errors.New("bad request")
	default:
		return nil, errors.New("unexpected response")
	}
}

func watchQuote(symbol string, frequency time.Duration, nasdaqChan chan NasdaqQuote) {
	defer close(nasdaqChan)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	ticker := time.NewTicker(frequency)
	defer ticker.Stop()

	// first request
	quote, err := requestQuote(symbol)
	if err != nil {
		log.Fatal(err)
		return
	}

	nasdaqChan <- *quote

	for {
		select {
		case <-signalChan:
			return
		case <-ticker.C:
			quote, err := requestQuote(symbol)
			if err != nil {
				log.Fatal(err)
				return
			}

			nasdaqChan <- *quote
		}
	}
}

func main() {
	symbol := flag.String("symbol", "AMD", "NASDAQ symbol")
	frequency := flag.Duration("frequency", 2*time.Second, "update frequency (i.e. 2s or 500ms)")
	flag.Parse()

	nasdaqChan := make(chan NasdaqQuote)

	go watchQuote(*symbol, *frequency, nasdaqChan)

	for n := range nasdaqChan {
		fmt.Printf("%s: %f\n", n.Symbol, n.Quote)
	}
}
