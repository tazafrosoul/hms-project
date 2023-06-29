package ginmux

import (
	s "hms-project/common/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddUser(in func(s.AddUserReq) s.AddUserRes) func(*gin.Context) {
	return func(ctx *gin.Context) {
		req := s.AddUserReq{}
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "could not bind JSON"})
		}

		user := in(req)

		//TODO implement checks and logs

		ctx.JSON(http.StatusOK, user)
	}
}
