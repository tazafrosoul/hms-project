package grpc

import (
	"context"
	"fmt"
	"hms-project/grpc/users/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type UserServer struct {
	pb.UnimplementedUserMgtServer
}

func NewUserServer() *UserServer {
	//TODO inject logs and other dependencies
	return &UserServer{}
}

func (s *UserServer) AddUser(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
	fmt.Printf("received from client the user to add: %s\n", in.FullName) //TODO remove this log line and inject add user logic

	return &pb.UserResponse{ //change this hardwired mess
		FullName: in.FullName,
		Username: in.Username,
		Email:    in.Email,
		Avatar:   in.Avatar,
		Role:     in.Role,
		ID:       "will be generated", //TODO generate id from database maybe.
	}, nil
	//TODO implement then inject the method logic here
}

func (s *UserServer) Run(addr string) {

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen to port %s : %v", addr, err)
	}

	registrar := grpc.NewServer()
	pb.RegisterUserMgtServer(registrar, s)

	if err := registrar.Serve(lis); err != nil {
		log.Fatalf("failed to listen to addr %s : %v", addr, err)
	}

}
