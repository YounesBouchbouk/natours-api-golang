package api

import (
	"net/http"

	db "github.com/YounesBouchbouk/natours-api-golang/db/sqlc"
	"github.com/YounesBouchbouk/natours-api-golang/token"
	"github.com/YounesBouchbouk/natours-api-golang/util"
	"github.com/gin-gonic/gin"
)

type CreateUserRequest struct {
	Name            string `json:"" binding:"required" validate:"required,min=3,max=100"`
	Email           string `json:"email" binding:"required" validate:"required,email"`
	Role            string `json:"role" binding:"required"`
	Photo           string `json:"photo"`
	Password        string `json:"password"`
	Confirmpassword string `json:"confirmpassword"`
}

func (server *Server) CreateUser(ctx *gin.Context) {

	var req CreateUserRequest

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	// check if passord matches checkpassword
	if req.Password != req.Confirmpassword {
		ctx.JSON(400, gin.H{
			"error": "passwords do not match",
		})
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	args := db.CreateUserParams{
		Email:    req.Email,
		Password: hashedPassword,
		Name:     req.Name,
		Role:     req.Role,
		Photo:    req.Photo,
	}

	user, err := server.store.CreateUser(ctx, args)

	user.Password = ""

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, user)
}

type GetUserRequest struct {
}

func (server *Server) getAllUsersForAdmin(ctx *gin.Context) {

	limit := int32(10)

	users, err := server.store.GetAllUsers(ctx, limit)

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": users,
	})
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required" validate:"required,email"`
	Password string `json:"password" binding:"required" validate:"required"`
}

func (server *Server) login(ctx *gin.Context) {

	maker, err := token.NewJWTMaker(server.config.JWT_secret_key)

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	var req LoginRequest

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := server.store.GetUserByEmail(ctx, req.Email)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if !util.CheckPasswordHash(req.Password, user.Password) {
		ctx.JSON(400, gin.H{
			"error": "user or passwords do not match",
		})
		return
	}

	token, payload, err := maker.CreateToken(user.Email, user.Role, 2000)

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "login success",
		// "user":    user,
		"payload": payload,
		"token":   token,
	})

}
