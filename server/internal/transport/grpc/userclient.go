package grpc

import (
	"context"
	"fmt"
	s "hms-project/common/structs"
	"hms-project/common/utility"
	"hms-project/grpc/users/pb"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	//TODO these vars are to be parsed as flags or as env variables
	address string = "localhost:8081"
)

type GrpcTrans struct{}

func NewGrpcTrans() *GrpcTrans {
	return &GrpcTrans{}
}

func (t *GrpcTrans) AddUser(aurq s.AddUserReq) (s.AddUserRes, error) {

	//TODO make this a separate GrpcTrans method
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not dial the grpc service: %v", err)
	}
	defer conn.Close()

	c := pb.NewUserMgtClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	hashedpw, err := utility.Hash(aurq.Password)
	if err != nil {
		//TODO add log service
		log.Fatalf("could not hash the provided password: %v\n", err)
		return s.AddUserRes{}, err
	}

	//structure mock user here------
	user := pb.UserRequest{
		By:       aurq.By,
		FullName: aurq.FullName,
		Username: aurq.Username,
		Email:    aurq.Email,
		Avatar:   aurq.Avatar,
		Password: hashedpw,
		Role:     aurq.Role,
	}

	//call remote procedure here------
	res, err := c.AddUser(ctx, &user)
	if err != nil {
		log.Fatalf("could not add user: %v", err)
	}
	fmt.Printf("received data from server: %v\n", res)
	return s.AddUserRes{
		ID:       res.ID,
		FullName: res.FullName,
		Username: res.Username,
		Email:    res.Email,
		Avatar:   res.Avatar,
		Role:     res.Role,
	}, nil
}
