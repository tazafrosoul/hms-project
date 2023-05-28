package main

import (
	"hms-project/server/internal/api"
	"hms-project/server/internal/api/mux/ginmux"
)

//TODO implement api transport layer

func main() {
	var (
		address string = ":5300"
	)

	//=========================================================================
	// service:=new(gRPC)
	// s:=newserver(service)
	// s.Run(address)

	//
	//pick up Http traffic from clients
	//Deliver the traffic to various Microservices through gRPC

	// s := gin.Default()
	// s.GET("/ping", TstHandler)

	// s.Run(":8080")
	//
	//========================================================================

	mux := ginmux.NewGin() //you can interchange with gorilla mux or any multiplexer
	api := api.NewApi(mux)
	api.Route()
	api.Run(address)

}
