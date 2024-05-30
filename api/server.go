package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/reflection/frog_blossom_db/db/sqlc"
)

// Server serves HTTP requets for CMS
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer creates new HTTP server and sets up routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	subrouter := router.Group("api/v1")
	subrouter.POST("/users", server.createUsers)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
