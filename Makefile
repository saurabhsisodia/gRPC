.PHONY: protos

protos:
	protoc -I protos/ protos/currency/currency.proto --go-grpc_out=protos/currency