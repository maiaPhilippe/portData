package handlers

import (
	"clientAPI/grpc"
	"clientAPI/protobuf"
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

func GetById(w http.ResponseWriter, r *http.Request) {

	cc := grpc.StartGrpcServer()
	client := protobuf.NewPortServiceClient(cc)
	defer cc.Close()

	id := strings.TrimPrefix(r.URL.Path, "/id/")

	request := &protobuf.GetPortByIdRequest{
		Key: id,
	}

	resp, _ := client.GetPortById(context.Background(), request)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}