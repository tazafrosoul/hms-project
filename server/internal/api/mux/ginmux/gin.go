package ginmux

import (
	i "hms-project/server/common/interfaces"
	s "hms-project/server/common/structs"
	"net/http"

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
func (g *Gin) Route() { //TODO requires gRPC transport to users microservice

	g.server.POST("/user", func(ctx *gin.Context) {
		req := s.AddUserReq{}
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "could not bind JSON"})
		}

		user := g.api.AddUser(req.Name)

		//TODO implement add new user

		ctx.JSON(http.StatusOK, user)
	})
}

func (g *Gin) Run(addr string) error {
	return g.server.Run(addr)
}
