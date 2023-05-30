package service

import (
	i "hms-project/server/common/interfaces"
	s "hms-project/server/common/structs"
)

type Service struct {
	grpcTrans i.GrpcTransport
}

func NewService(gt i.GrpcTransport) *Service {
	return &Service{
		grpcTrans: gt,
	}
}

func (as *Service) AddUser(name string) s.User {
	//implement logic here
	return as.grpcTrans.AddUser(name)
}
