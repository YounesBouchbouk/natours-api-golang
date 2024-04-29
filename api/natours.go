package api

import (
	"github.com/gin-gonic/gin"
)

func (server *Server) privateroutetest(ctx *gin.Context) {
	email, _ := ctx.Get("email")
	role, _ := ctx.Get("role")

	ctx.JSON(200, gin.H{
		"email": email,
		"role":  role,
	})

}
