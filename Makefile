.PHONY=protos

protos:
	protoc --proto_path=./protos --go_out=plugins=grpc:./protos differenceProduct.proto