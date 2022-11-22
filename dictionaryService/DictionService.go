package main

import service "github.com/franklynobleC/dictionaryAPIGrpc/pb"

// import "google.golang.org/genproto/googleapis/cloud/orchestration/airflow/service/v1"

type server struct {
	service.UnimplementedEnglishDictionaryServer
}
