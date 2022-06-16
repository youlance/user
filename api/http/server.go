package server

import (
	"github.com/gin-gonic/gin"
	db "github.com/youlance/user/db/sqlc"
	"github.com/youlance/user/pkg/config"
)

type Server struct {
	config config.Config
	db     *db.Queries
	router *gin.Engine
}

func NewServer(config config.Config, db *db.Queries) (*Server, error) {
	server := &Server{
		config: config,
		db:     db,
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/user", server.CreateUser)
	router.GET("/user/:username", server.GetUser)

	router.POST("/follower", CreateUserFollower)
	router.GET("/follower/:username", GetUserFollower)
	router.GET("/followee/:username", GetUserFollowee)
	router.GET("/followers/:username", GetFollowers)
	router.GET("/followees/:username", GetFollowees)
	router.DELETE("/follower/:username", DeleteFollower)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}