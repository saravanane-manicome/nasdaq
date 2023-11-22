package service

type Quote struct {
	Symbol string
	Quote  float64
	Exists bool
}

type IQuoteService interface {
	GetQuote(string) (*Quote, error)
}
