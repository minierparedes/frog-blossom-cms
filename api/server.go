package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/reflection/frog_blossom_db/db/sqlc"
	"github.com/reflection/frog_blossom_db/internal/handler"
)

// Server serves HTTP requets for CMS
type Server struct {
	Store  *db.Store
	router *gin.Engine
}

// NewServer creates new HTTP server and sets up routing
func NewServer(store *db.Store) *Server {
	server := &Server{Store: store}
	router := gin.Default()

	subrouter := router.Group("api/v1")

	subrouter.POST("/users", handler.CreateUsersHandler(store))
	subrouter.GET("/users/:id", handler.GetUsersHandler(store))
	subrouter.GET("/users", handler.ListUsersHandler(store))
	subrouter.POST("/users/", handler.UpdateUsersHandler(store))

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
