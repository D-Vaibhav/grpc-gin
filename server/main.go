package main

import (
	"context"
	"fmt"
	"net"

	"github.com/vaibhav/grpc_gin/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
}

var Port = ":4000"

func main() {
	// to implement MicroserviceServer
	listner, err := net.Listen("tcp", Port)
	if err != nil {
		panic(err)
	}
	fmt.Println("Server start listenig to the Port", Port)

	grpcServer := grpc.NewServer()
	protos.RegisterSubtractDivideServer(grpcServer, &server{})

	// now for serializing and de-serializing data
	reflection.Register(grpcServer)

	// finally serving grpcServer
	err = grpcServer.Serve(listner)
	if err != nil {
		panic(err)
	}
}

// implementing the interface so to RegisterMicroserviceServer
func (s *server) CalculateDifference(ctx context.Context, req *protos.Request) (*protos.Response, error) {
	numberFirst, numberSecond := req.GetNumberFirst(), req.GetNumberSecond()

	result := numberFirst - numberSecond
	return &protos.Response{CalculatedAnswer: result}, nil
}
func (s *server) CalculateProduct(ctx context.Context, req *protos.Request) (*protos.Response, error) {
	numberFirst, numberSecond := req.GetNumberFirst(), req.GetNumberSecond()

	result := numberFirst * numberSecond
	return &protos.Response{CalculatedAnswer: result}, nil
}
