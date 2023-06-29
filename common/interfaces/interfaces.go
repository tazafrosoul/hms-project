package i

import s "hms-project/common/structs"

type Server interface {
	Route()
	Run(addr string) error
}

type ApiService interface {
	AddUser(s.AddUserReq) s.AddUserRes
}

type UserService interface {
	AddUser(s.AddUserReq) (s.AddUserRes, error)
}

type GrpcTransport interface {
	AddUser(s.AddUserReq) s.AddUserRes
}

type Repo interface {
	Create(s.User) (string, error)
}

type DB interface {
	Run()
}
