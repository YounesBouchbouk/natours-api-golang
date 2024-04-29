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

	router.Use(AuthenticationMiddlware(*server))

	//TODO

	// get all tours
	// get active tours
	// get tour by it's location
	// get tour information
	// get tour by defficulty
	// get tour from to

	// ONLY FOR ADMIN
	router.GET("/users", CheckIfAdmin(), server.getAllUsersForAdmin)

	// TODO

	// remove user /user/remove/{id}
	// check active or  disbaled user /users/status/{role}
	// get user by id /users/{id}
	// get users by role  /users/role/{role}

	router.GET("/testroute", server.privateroutetest)

	server.router = router

	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
