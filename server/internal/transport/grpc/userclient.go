package grpc

import (
	"context"
	"fmt"
	"log"
	"time"

	s "github.com/tazafrosoul/hms-project/common/structs"
	"github.com/tazafrosoul/hms-project/common/utility"
	"github.com/tazafrosoul/hms-project/grpc/users/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Transport struct {
	address string
}

func NewTransport(address string) *Transport {
	return &Transport{
		address: address,
	}
}

func (t *Transport) AddUser(aurq s.AddUserReq) (s.AddUserRes, error) {

	//Dialing GrPC server
	conn, err := grpc.Dial("localhost"+t.address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not dial the grpc service: %v", err)
	}
	defer conn.Close()

	//Pushing GrPC request
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
		Avatar:   res.Avatar,
		Role:     res.Role,
	}, nil
}
