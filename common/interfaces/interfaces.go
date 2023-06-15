package i

import s "hms-project/common/structs"

type Server interface {
	Route()
	Run(addr string) error
}

type ApiService interface {
	AddUser(s.AddUserReq) s.AddUserRes
}

type GrpcTransport interface {
	AddUser(s.AddUserReq) s.AddUserRes
}
