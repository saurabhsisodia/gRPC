package currency
import (
	"log"
	"context"

	data "github.com/saurabhsisodia/gRPC/protos/data"
)

type Currency struct{
	logger *log.Logger
	cr *data.CurrencyRate
}

func NewCurrency(l *log.Logger)*Currency{
	cr,err:=data.NewCurrencyRate(l)
	if err!=nil{
		log.Fatal(err)
	}
	return &Currency{l,cr}
}

func (c *Currency) GetRate(ctx context.Context,rr *RateRequest) (*RateResponse,error){
	c.logger.Println("Handle GetRate","base",rr.GetBase(),"Destination",rr.GetDestination())

	rate,err:=c.cr.GetRate(rr.GetBase().String(),rr.GetDestination().String())
	if err!=nil{
		c.logger.Fatal(err)
	}
	return &RateResponse{Rate:rate},nil

}

func (c *Currency) mustEmbedUnimplementedCurrencyServer(){}

