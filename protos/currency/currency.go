package currency
import (
	"log"
	"context"
)

type Currency struct{
	logger *log.Logger
}

func NewCurrency(l *log.Logger)*Currency{
	return &Currency{l}
}

func (c *Currency) GetRate(ctx context.Context,rr *RateRequest) (*RateResponse,error){
	c.logger.Println("Handle GetRate","base",rr.GetBase(),"Destination",rr.GetDestination())
	return &RateResponse{Rate:0.5},nil

}

func (c *Currency) mustEmbedUnimplementedCurrencyServer(){}

