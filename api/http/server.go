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

	router.Use(CORS())

	router.POST("/user", server.CreateUser)
	router.GET("/user/:username", server.GetUser)
	router.POST("/user/verify", server.VerifyUser)

	router.POST("/follower", server.CreateUserFollower)
	router.GET("/followers", server.ListFollowers)
	router.GET("/followees", server.ListFollowees)
	router.GET("/followees/count/:username", server.GetUserFolloweesCount)
	router.GET("/followers/count/:username", server.GetUserFollowersCount)
	router.DELETE("/follower", server.DeleteFollower)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
