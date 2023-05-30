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
	fmt.Printf("received from client the username: %s\n", in.Name)
	return &pb.UserResponse{
		Name: in.Name,
		Role: "test",
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
