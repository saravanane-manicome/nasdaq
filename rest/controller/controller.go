package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/saravanane-manicome/nasdaq/rest/service"
	"log"
	"net/http"
)

type quoteQuery struct {
	Symbol string `form:"symbol" binding:"required"`
}

type QuoteResponse struct {
	Symbol string
	Quote  float64
}

type QuoteController struct {
	QuoteService service.IQuoteService
}

func (controller *QuoteController) Serve() {
	r := gin.Default()

	r.GET("/quote", controller.getQuote)

	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func (controller *QuoteController) getQuote(context *gin.Context) {
	var q quoteQuery
	if err := context.ShouldBindWith(&q, binding.Query); err != nil {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	quote, err := controller.QuoteService.GetQuote(q.Symbol)

	if err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if quote == nil {
		context.AbortWithStatus(http.StatusNotFound)
		return
	}

	response := QuoteResponse{quote.Symbol, quote.Quote}

	context.JSON(http.StatusOK, response)
}
