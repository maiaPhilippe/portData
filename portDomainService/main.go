package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"portDomainService/handlers"
	"portDomainService/protobuf"
)

type Server struct {
}

func (s Server) PortData(ctx context.Context, request *protobuf.PortRequest) (*protobuf.PortResponse, error) {
	return handlers.InsertPortData(ctx, request)
}

func (s Server) GetPortById(ctx context.Context, request *protobuf.GetPortByIdRequest) (*protobuf.GetPortByIdResponse, error) {
	return handlers.GetById(ctx, request)
}

func main() {

	address := ":9090"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	fmt.Printf("Server is listening on %v ...", address)

	s := grpc.NewServer()
	protobuf.RegisterPortServiceServer(s, &Server{})
	s.Serve(lis)
}

