package service

import (
	i "hms-project/common/interfaces"
	s "hms-project/common/structs"
	"log"
)

type Service struct {
	grpcTrans i.GrpcTransport
}

func NewService(gt i.GrpcTransport) *Service {
	return &Service{
		grpcTrans: gt,
	}
}

func (ss *Service) AddUser(aurq s.AddUserReq) (s.AddUserRes, error) {
	//implement logic here
	result, err := ss.grpcTrans.AddUser(aurq)
	if err != nil {
		//TODO implement log here
		log.Fatalf("experienced error: %v\n", err)
		return s.AddUserRes{}, err
	}
	return result, nil
}
