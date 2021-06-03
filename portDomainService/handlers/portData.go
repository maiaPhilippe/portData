package handlers

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"portDomainService/protobuf"
	"portDomainService/server"
)



func InsertPortData(ctx context.Context, request *protobuf.PortRequest) (*protobuf.PortResponse, error) {

	client := server.ConnectDatabase()
	defer client.Disconnect(ctx)

	portDatabase := client.Database("port_data")
	portCollection := portDatabase.Collection("port")

	opts := options.UpdateOptions{}
	upsertBool := true
	opts.Upsert = &upsertBool

	_, err := portCollection.UpdateOne(ctx, bson.M{"key": request.Key}, bson.M{"$set": *request}, &opts)
	if err != nil {
		return nil, err
	}

	response := &protobuf.PortResponse{
		Response: "updated",
	}
	return response, nil
}

