package quote

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type PriceRequestModel struct {
	Age            *int      `json:"age" binding:"required"`
	CoverStartDate time.Time `json:"coverStartDate" binding:"required"`
}

type PriceResponseModel struct {
	Price     float64 `json:"price"`
	TaxAmount float64 `json:"taxAmount"`
	QuoteId   string  `json:"quoteId"`
}

// Create an interface so we can create an instance to use in main()
type QuoteController interface {
	HandleQuotePostRequest(c *gin.Context)
}

// This is the underlying type that this conroller will end up sitting under
type quoteController struct {
}

// Constructor
func NewQuoteController() QuoteController {
	return &quoteController{}
}

// Function is passed an instance of the quoteController so is a method attached to that struct
func (s *quoteController) HandleQuotePostRequest(c *gin.Context) {
	var req PriceRequestModel
	err := c.BindJSON(&req)
	if err != nil {
		// Obviously need to handle/trap errors better
		return
	}

	netPrice := float64(*req.Age) * 2
	tax := netPrice * 0.2
	grossPrice := netPrice + tax

	res := PriceResponseModel{
		Price:     grossPrice,
		TaxAmount: tax,
		QuoteId:   "QID" + strconv.Itoa(rand.Intn(1000)),
	}

	c.IndentedJSON(http.StatusOK, res)
}
