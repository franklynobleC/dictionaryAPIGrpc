protos: 
	 protoc --proto_path=pb --go_out=pb --go_opt=paths=source_relative \
	 --go-grpc_out=pb  --go-grpc_opt=paths=source_relative \
	 --grpc-gateway_out ./pb  --grpc-gateway_opt paths=source_relative \
	 pb/*.proto

tools:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2 
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway 
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 


pulsar:
	 GOLANG_VERSION=1.18 PULSAR_VERSION=2.10.0


# gos:
#   go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway 
#   go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 




# doc: 
 
#   docker run -it -p 6650:6650  -p 8080:8080 --mount source=pulsardata,target=/pulsar/data --mount source=pulsarconf,target=/pulsar/conf apachepulsar/pulsar:@pulsar:version@ bin/pulsar standalone	 