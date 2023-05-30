package i

import s "hms-project/server/common/structs"

type Server interface {
	Route()
	Run(addr string) error
}

type ApiService interface {
	AddUser(string) s.User
}

type GrpcTransport interface {
	AddUser(string) s.User
}
