package handlers

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"portDomainService/protobuf"
	"portDomainService/server"
)


func GetById(ctx context.Context, request *protobuf.GetPortByIdRequest) (*protobuf.GetPortByIdResponse, error) {

	client := server.ConnectDatabase()
	defer client.Disconnect(ctx)

	portDatabase := client.Database("port_data")
	portCollection := portDatabase.Collection("port")

	opts := options.UpdateOptions{}
	upsertBool := true
	opts.Upsert = &upsertBool

	var portData protobuf.GetPortByIdResponse
	err := portCollection.FindOne(ctx, bson.M{"key": request.Key}).Decode(&portData)
	if err != nil {
		return nil, err
	}

	response := &protobuf.GetPortByIdResponse{
		Name: portData.Name,
		City: portData.City,
		Country: portData.Country,
		Alias: portData.Alias,
		Regions: portData.Regions,
		Coordinates: portData.Coordinates,
		Province: portData.Province,
		Timezone: portData.Timezone,
		Unlocs: portData.Unlocs,
		Code: portData.Code,
	}
	return response, nil
}

