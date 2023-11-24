package service

type Quote struct {
	Symbol string
	Quote  float64
}

type IQuoteService interface {
	GetQuote(string) (*Quote, error)
}
