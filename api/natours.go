package api

import (
	"github.com/YounesBouchbouk/natours-api-golang/token"
	"github.com/gin-gonic/gin"
)

func (server *Server) privateroutetest(ctx *gin.Context) {
	authpayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	ctx.JSON(200, gin.H{
		"email": authpayload.Email,
		"role":  authpayload.Role,
	})

}
