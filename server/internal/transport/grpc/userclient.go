package grpc

import (
	"context"
	"fmt"
	s "hms-project/common/structs"
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

func (t *GrpcTrans) AddUser(name string) s.User {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not dial the grpc service: %v", err)
	}
	defer conn.Close()

	c := pb.NewUserMgtClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//structure mock user here------
	user := pb.UserRequest{
		Name: name,
	}

	//call remote procedure here------
	res, err := c.AddUser(ctx, &user)
	if err != nil {
		log.Fatalf("could not add user: %v", err)
	}
	fmt.Printf("received data from server: %v\n", res)
	return s.User{
		Name: res.Name,
		Role: res.Role,
	}
}
