package ginmux

import "github.com/gin-gonic/gin"

type Gin struct {
	server *gin.Engine
}

func NewGin() *Gin {
	s := gin.Default()
	return &Gin{
		server: s,
	}
}

// ---------ROUTES---------
func (g *Gin) Route() {
	g.server.POST("/user", HandleAddUser) //TODO requires gRPC transport to users microservice
}

func (g *Gin) Run(addr string) error {
	return g.server.Run(addr)
}
