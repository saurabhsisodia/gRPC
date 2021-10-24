package data

import (
	"testing"
	"fmt"
	"log"
	"os"
)

func TestCurrencyData(t *testing.T){

	l:=log.New(os.Stdout,"[INFO] Rate-Exchange-Data: ",log.LstdFlags)

	cr,err:=NewCurrencyRate(l)

	if err!=nil{
		t.Fatal(err)
	}
	fmt.Printf("%#v\n",cr.rate)

}