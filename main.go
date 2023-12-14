package main

import (
	"simple-quote-api-go/controllers/quote"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	quoteController := quote.NewQuoteController()
	router.POST("/quote", quoteController.HandleQuotePostRequest)

	router.Run("localhost:8080")
}
