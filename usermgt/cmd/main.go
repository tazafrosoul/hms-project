package main

import (
	"hms-project/usermgt/internal/repository"
	"hms-project/usermgt/internal/transport/grpc"
	"hms-project/usermgt/internal/userservice"
)

var (
	address string = ":8081"
)

func main() {
	//storage repo
	repository := repository.NewMemoRepo() //chose inmemory. can swap with orchestra (mix of redis and postgres)

	//service layer. this is a linkage
	userservice := userservice.NewUserService(repository) //injected repository

	// transpor layer. connects this microservice with other microservices. can have multiple of these, other using
	// Websockets, rabbitmq etc. all being instances of multithread go routines
	server := grpc.NewUserTransport(userservice) //chose gRPC
	server.Run(address)
}
