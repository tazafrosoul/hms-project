package ginmux

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleAddUser(ctx *gin.Context) {
	//TODO implement add new user
	ctx.JSON(http.StatusOK, gin.H{"TODO": "handle adding new user"})
}
