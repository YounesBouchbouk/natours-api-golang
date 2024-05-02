package api

import (
	"fmt"

	db "github.com/YounesBouchbouk/natours-api-golang/db/sqlc"
	"github.com/YounesBouchbouk/natours-api-golang/token"
	"github.com/YounesBouchbouk/natours-api-golang/util"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router     *gin.Engine
	store      *db.Store
	config     *util.Config
	tokenMaker token.Maker
}

func NewServer(store *db.Store, config *util.Config) (*Server, error) {
	maker, err := token.NewJWTMaker(config.JWT_secret_key)

	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		store:      store,
		config:     config,
		tokenMaker: maker,
	}
	router := gin.Default()

	//add router
	router.POST("/user", server.CreateUser)
	router.POST("/login", server.login)

	router.POST("/tour", server.createNewTourController)

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

	return server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
