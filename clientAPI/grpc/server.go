package grpc

import (
	"google.golang.org/grpc"
	"log"
)

func StartGrpcServer() *grpc.ClientConn{
	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("port-service:9090", opts)
	if err != nil {
		log.Fatal(err)
	}

	return cc
}
