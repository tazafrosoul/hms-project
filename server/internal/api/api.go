package api

import (
	"log"

	i "github.com/tazafrosoul/hms-project/common/interfaces"
)

type Api struct {
	m i.Server
}

// instantiate Api struct (accept any multiplexer due to [i.Server] interface)
func NewApi(mux i.Server) *Api {
	return &Api{
		m: mux,
	}
}

func (a *Api) Route() {
	a.m.Route()
}

func (a *Api) Run(addr string) error {
	if err := a.m.Run(addr); err != nil {
		log.Fatalf("failed to serve on addr %s :: %s", addr, err.Error())
		return err
	}
	return nil
}
