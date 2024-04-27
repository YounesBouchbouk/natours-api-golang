package api

import (
	db "github.com/YounesBouchbouk/natours-api-golang/db/sqlc"
	"github.com/YounesBouchbouk/natours-api-golang/util"
	"github.com/gin-gonic/gin"
)

func (server *Server) createUser(ctx *gin.Context) {

	args := db.CreateUserParams{
		Email:    util.RandomString(20) + "@gmail.com",
		Password: util.RandomString(10),
		Name:     util.RandomString(10),
	}

	user, err := server.store.CreateUser(ctx, args)

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, user)
}
