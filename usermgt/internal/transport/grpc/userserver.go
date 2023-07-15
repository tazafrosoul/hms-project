package grpc

import (
	"context"
	"fmt"
	i "hms-project/common/interfaces"
	s "hms-project/common/structs"
	"hms-project/grpc/users/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type UserTransport struct {
	usi i.UserService
	pb.UnimplementedUserMgtServer
}

func NewUserTransport(usi i.UserService) *UserTransport {
	//TODO inject log
	return &UserTransport{
		usi: usi,
	}
}

func (ut *UserTransport) AddUser(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
	fmt.Printf("received from client the user: %s\n", in.FullName)

	request := s.AddUserReq{
		By:       in.By,
		FullName: in.FullName,
		Username: in.Username,
		Email:    in.Email,
		Avatar:   in.Avatar,
		Role:     in.Role,
		Password: in.Password, //TODO hash password + validate on the http server
	}

	res, err := ut.usi.AddUser(request)
	if err != nil {
		return &pb.UserResponse{}, err
	}

	return &pb.UserResponse{ //change this hardwired mess
		FullName: res.FullName,
		Username: res.Username,
		Email:    res.Email,
		Avatar:   res.Avatar,
		Role:     res.Role,
		ID:       res.ID,
	}, nil
}

func (ut *UserTransport) Run(addr string) {

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen to port %s : %v", addr, err)
	}

	registrar := grpc.NewServer()
	pb.RegisterUserMgtServer(registrar, ut)

	if err := registrar.Serve(lis); err != nil {
		log.Fatalf("failed to listen to addr %s : %v", addr, err)
	}

}
