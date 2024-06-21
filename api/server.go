package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/reflection/frog-blossom-cms/db/sqlc"
	"github.com/reflection/frog-blossom-cms/internal/handler"
)

// Server serves HTTP request for CMS
type Server struct {
	Store  db.Store
	router *gin.Engine
}

// NewServer creates new HTTP server and sets up routing
func NewServer(store db.Store) *Server {
	server := &Server{Store: store}
	router := gin.Default()

	subrouter := router.Group("api/v1")

	// Users router
	subrouter.POST("/users", handler.CreateUsersHandler(store))
	subrouter.PUT("/users/:id", handler.UpdateUserHandler(store))
	subrouter.GET("/users/:id", handler.GetUsersHandler(store))
	subrouter.GET("/users", handler.ListUsersHandler(store))
	subrouter.DELETE("/users/:id", handler.SoftDeleteUsersHandler(store))

	// Pages router
	subrouter.POST("/pages", handler.CreatePageTxHandler(store))
	subrouter.PUT("/pages/:id", handler.UpdatePagesTxHandler(store))
	subrouter.GET("/pages/:id", handler.GetPageHandler(store))
	subrouter.GET("/pages", handler.ListPagesHandler(store))
	subrouter.DELETE("/pages/:id", handler.DeletePageTxHandler(store))

	// Posts router
	subrouter.POST("/posts", handler.CreatePostTxHandler(store))
	subrouter.GET("/posts/:id", handler.GetPostHandler(store))
	
	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
