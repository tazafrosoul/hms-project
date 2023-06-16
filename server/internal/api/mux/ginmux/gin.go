package ginmux

import (
	i "hms-project/common/interfaces"

	"github.com/gin-gonic/gin"
)

type Gin struct {
	server *gin.Engine
	api    i.ApiService
}

func NewGin(service i.ApiService) *Gin {
	s := gin.Default()
	return &Gin{
		server: s,
		api:    service,
	}
}

// ---------ROUTES---------
func (g *Gin) Route() {

	g.server.POST("/user", AddUser(g.api.AddUser))
}

func (g *Gin) Run(addr string) error {
	return g.server.Run(addr)
}
