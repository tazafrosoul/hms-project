package i

import s "github.com/tazafrosoul/hms-project/common/structs"

type Server interface {
	Route()
	Run(addr string) error
}

type ApiService interface {
	AddUser(s.AddUserReq) (s.AddUserRes, error)
}

type UserService interface {
	AddUser(s.AddUserReq) (s.AddUserRes, error)
}

type GrpcTransport interface {
	AddUser(s.AddUserReq) (s.AddUserRes, error)
}

type Repo interface {
	Create(s.User) (string, error)
}

type DB interface {
	Run()
}

type Core interface {
	Hash(string) string
	validate(s.User) bool
}
