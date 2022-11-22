protos: 
	 protoc --proto_path=pb --go_out=pb --go_opt=paths=source_relative \
	 --go-grpc_out=pb  --go-grpc_opt=paths=source_relative \
	 pb/*.proto

tools:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2 

	
pulsar:
	 GOLANG_VERSION=1.18 PULSAR_VERSION=2.10.0