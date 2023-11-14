package main

import (
	"os"

	"github.com/tazafrosoul/hms-project/usermgt/internal/repository"
	"github.com/tazafrosoul/hms-project/usermgt/internal/transport/grpc"
	"github.com/tazafrosoul/hms-project/usermgt/internal/userservice"
)

func main() {
	//Load environment variables
	// utility.Loadvars(".env.dev")

	var (
		grpcaddress string = os.Getenv("GRPC_SERVER_ADDRESS")
	)

	//storage repo
	repository := repository.NewMemoRepo() //chose inmemory. can swap with orchestra (mix of redis and postgres)
	repository.Init()

	//service layer. this is a linkage
	userservice := userservice.NewUserService(repository) //injected repository

	// transpor layer. connects this microservice with other microservices. can have multiple of these, other using
	// Websockets, rabbitmq etc. all being instances of multithread go routines
	server := grpc.NewUserTransport(userservice) //chose gRPC
	server.Run(grpcaddress)
}
