proto:
	protoc --proto_path=proto --go_out=plugins=grpc:portDomainService --go_out=plugins=grpc:clientAPI port.proto

run_port_server:
	cd ./portDomainService && go run main.go

run_client:
	cd ./clientAPI && go run main.go

mongo:
	docker run --name port_mongodb -p 27017:27017 -d mongo:latest