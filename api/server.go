package api

import (
	db "github.com/YounesBouchbouk/natours-api-golang/db/sqlc"
	"github.com/YounesBouchbouk/natours-api-golang/util"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	store  *db.Store
	config *util.Config
}

func NewServer(store *db.Store, config *util.Config) *Server {
	server := &Server{
		store:  store,
		config: config,
	}
	router := gin.Default()

	//add router
	router.POST("/user", server.CreateUser)
	router.POST("/login", server.login)

	//admin route
	router.GET("/users", server.getAllUsersForAdmin)

	server.router = router

	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
