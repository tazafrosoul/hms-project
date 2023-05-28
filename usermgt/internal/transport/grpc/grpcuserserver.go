package grpc

import (
	"context"
	"hms-project/usermgt/internal/transport/grpc/usersgrpc"
	"log"
	"net"

	"google.golang.org/grpc"
)

type UserServer struct {
	usersgrpc.UnimplementedUserMgtServer
}

func NewUserServer() *UserServer {
	//TODO this will have dependency injection
	return &UserServer{}
}

func (s *UserServer) GetUser(context.Context, *usersgrpc.UserRequest) (*usersgrpc.UserResponse, error) {
	return &usersgrpc.UserResponse{
		Name: "test",
		Role: "test",
	}, nil
	//TODO implement the method logic here
}

func (s *UserServer) Run(addr string) {

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("cannot listen to said port : %v", err)
	}

	registrar := grpc.NewServer()
	usersgrpc.RegisterUserMgtServer(registrar, s)

	registrar.Serve(lis)
}
