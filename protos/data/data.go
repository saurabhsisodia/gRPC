package data

import (
	"log"
	"net/http"
	"strconv"
	"fmt"
	"encoding/xml"
	
)

type CurrencyRate struct{
	l *log.Logger
	rate map[string]float64
}

// constructor
func (cr *CurrencyRate) GetRate(base,destination string) (float64,error){
	baseRate,ok:=cr.rate[base]
	if !ok{
		return 0,fmt.Errorf("Invalid base %s\n",base)
	}
	destRate,ok:=cr.rate[destination]
	if !ok{
		return 0,fmt.Errorf("Invalid Destination %s\n",destination)
	}

	return destRate/baseRate,nil
}
func NewCurrencyRate (l *log.Logger)(*CurrencyRate,error){
	mp:=make(map[string]float64)
	cr:=&CurrencyRate{l,mp}

	err:=cr.GetCurrencyRate()
	return cr,err
}

func (cr *CurrencyRate) GetCurrencyRate() error{
	resp,err:=http.DefaultClient.Get("https://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml?c2f385ad549b384828d2bf5c643110bd")
	if err!=nil{
		return err
	}

	if resp.StatusCode!=http.StatusOK{
		return fmt.Errorf("Got StatusCode %d\n",resp.StatusCode)
	}

	defer resp.Body.Close()

	rd:=&Cubes{}
	err = xml.NewDecoder(resp.Body).Decode(rd)
	if err!=nil{
		return err
	}

	for _,c:=range rd.CubeData{
		r,err:=strconv.ParseFloat(c.Rate,64)
		if err!=nil{
			return err
		}
		cr.rate[c.Currency]=r
	}
	return nil
}

type Cubes struct{
	CubeData []Cube `xml:"Cube>Cube>Cube"`
}
type Cube struct{
	Currency string `xml:"currency,attr"`
	Rate string `xml:"rate,attr"`
}