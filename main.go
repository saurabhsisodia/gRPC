/*
	protoc --go_out=. --go_opt=paths=source_relative currency.proto
	protoc --go-grpc_out=. --go-grpc_opt=paths=source_relative currency.proto
*/
package main

import (
	"log"
	"os"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	currency "github.com/saurabhsisodia/gRPC/protos/currency"
)
func main(){
	l:=log.New(os.Stdout,"gRPC [INFO] ",log.LstdFlags)
	cs:=currency.NewCurrency(l)
	gs:=grpc.NewServer()

	currency.RegisterCurrencyServer(gs,cs)
	reflection.Register(gs)   // remove in production
	listener,err:=net.Listen("tcp",":9092")
	if err!=nil{
		log.Fatalf("Unable to Listen on port 9092 %v\n",err)
	}
	gs.Serve(listener)
}