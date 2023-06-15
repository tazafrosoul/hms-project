package service

import (
	i "hms-project/common/interfaces"
	s "hms-project/common/structs"
)

type Service struct {
	grpcTrans i.GrpcTransport
}

func NewService(gt i.GrpcTransport) *Service {
	return &Service{
		grpcTrans: gt,
	}
}

func (ss *Service) AddUser(aurq s.AddUserReq) s.AddUserRes {
	//implement logic here
	return ss.grpcTrans.AddUser(aurq)
}
