package main

import (
	"os"

	"github.com/tazafrosoul/hms-project/server/internal/api"
	"github.com/tazafrosoul/hms-project/server/internal/api/mux/ginmux"
	"github.com/tazafrosoul/hms-project/server/internal/service"
	"github.com/tazafrosoul/hms-project/server/internal/transport/grpc"
)

//TODO implement api transport layer

func main() {
	// utility.Loadvars("../.env.dev") //this will compile within docker image. Will not load environment if built here.

	var (
		httpaddress string = os.Getenv("HTTP_SERVER_PORT")
		grpcaddress string = os.Getenv("GRPC_SERVER_ADDRESS")
	)

	gt := grpc.NewTransport(grpcaddress)
	srv := service.NewService(gt)
	mux := ginmux.NewGin(srv) //you can interchange with gorilla mux or any multiplexer
	api := api.NewApi(mux)
	api.Route()
	api.Run(httpaddress)
}
