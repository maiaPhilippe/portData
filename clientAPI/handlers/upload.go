package handlers

import (
	"bytes"
	"clientAPI/grpc"
	"clientAPI/models"
	"clientAPI/protobuf"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

const MAX_MEMORY = 1 * 1024 * 1024

func readFile(w http.ResponseWriter, r *http.Request) (*bytes.Buffer, error)  {
	if err := r.ParseMultipartForm(MAX_MEMORY); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusForbidden)
		return nil, err
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return nil, err
	}
	defer file.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		log.Println(err)
	}

	return buf, nil
}

func parser(buf *bytes.Buffer) map[string]models.JsonModel{

	var ports map[string]models.JsonModel

	err := json.Unmarshal(buf.Bytes(), &ports)
	if err != nil {
		log.Fatal(err)
	}

	return ports
}

func Upload(w http.ResponseWriter, r *http.Request) {

	cc := grpc.StartGrpcServer()
	defer cc.Close()

	bufferFile, err := readFile(w, r)
	if err != nil {
		return
	}

	ports := parser(bufferFile)

	client := protobuf.NewPortServiceClient(cc)

	for key, v := range ports {
		request := &protobuf.PortRequest{
			Name: v.Name,
			City: v.City,
			Country: v.Country,
			Alias: v.Alias,
			Regions: v.Regions,
			Coordinates: v.Coordinates,
			Province: v.Province,
			Timezone: v.Timezone,
			Unlocs: v.Unlocs,
			Code: v.Code,
			Key: key,
		}

		client.PortData(context.Background(), request)
		//fmt.Printf("Receive response => [%v]", resp.Response)

	}
}
