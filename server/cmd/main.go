package main

import (
	"hms-project/server/internal/api"
	"hms-project/server/internal/api/mux/ginmux"
	"hms-project/server/internal/service"
	"hms-project/server/internal/transport/grpc"
)

//TODO implement api transport layer

func main() {
	var (
		address string = ":5300"
	)

	gt := grpc.NewGrpcTrans()
	srv := service.NewService(gt)
	mux := ginmux.NewGin(srv) //you can interchange with gorilla mux or any multiplexer
	api := api.NewApi(mux)
	api.Route()
	api.Run(address)

}
