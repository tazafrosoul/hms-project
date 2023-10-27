package ginmux

import (
	"net/http"

	s "github.com/tazafrosoul/hms-project/common/structs"

	"github.com/gin-gonic/gin"
)

func AddUser(in func(s.AddUserReq) (s.AddUserRes, error)) func(*gin.Context) {
	return func(ctx *gin.Context) {
		req := s.AddUserReq{}
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "could not bind JSON"})
		}

		res, err := in(req)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			//TODO implement checks and logs

			ctx.JSON(http.StatusOK, res)
		}

	}
}
