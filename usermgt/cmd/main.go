package main

import "hms-project/usermgt/internal/transport/grpc"

var (
	address string = ":8081"
)

func main() {
	server := grpc.NewUserServer()
	server.Run(address)

}
