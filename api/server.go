package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/reflection/frog-blossom-cms/common"
	"github.com/reflection/frog-blossom-cms/config"
	db "github.com/reflection/frog-blossom-cms/db/sqlc"
	"github.com/reflection/frog-blossom-cms/docs"
	"github.com/reflection/frog-blossom-cms/internal/handler"
	"github.com/reflection/frog-blossom-cms/internal/middleware"
	"github.com/reflection/frog-blossom-cms/internal/roles"
	"github.com/reflection/frog-blossom-cms/token"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title frog blossom API documentation
// @version 1
// @Description frog-blossom

// @host localhost:8080
// @BasePath /api/v1

// NewServer creates new HTTP server and sets up routing
func NewServer(config config.Config, store db.Store) (*common.Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.TokenSystemmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &common.Server{
		Store:      store,
		TokenMaker: tokenMaker,
		Config:     config,
	}
	router := gin.Default()

	docs.SwaggerInfo.BasePath = "/api/v1"

	subrouter := router.Group("api/v1")
	// authRoutes authorization middleware
	// authRoutes := router.Group("/").Use(middleware.AuthMiddleware(server.TokenMaker))

	// User login
	subrouter.POST("/users/login", handler.LoginUserHandler(server, store))

	// Users router
	subrouter.POST("/users", middleware.Authentication(server.TokenMaker), middleware.Authorization(roles.CreateUser), handler.CreateUsersHandler(store))
	subrouter.POST("/users/register", middleware.Authentication(server.TokenMaker), middleware.Authorization(roles.CreateUser), handler.CreateInitialAdminHandler(store))
	subrouter.PUT("/users/:id", middleware.Authentication(server.TokenMaker), middleware.Authorization(roles.UpdateUser), handler.UpdateUserHandler(store))
	subrouter.GET("/users/:id", middleware.Authentication(server.TokenMaker), middleware.Authorization(roles.GetUser), handler.GetUsersHandler(store))
	subrouter.GET("/users", middleware.Authentication(server.TokenMaker), middleware.Authorization(roles.ListUsers), handler.ListUsersHandler(store))
	subrouter.DELETE("/users/:id", middleware.Authentication(server.TokenMaker), middleware.Authorization(roles.DeleteUser), handler.SoftDeleteUsersHandler(store))

	// Pages router
	subrouter.POST("/pages", middleware.Authentication(server.TokenMaker), middleware.Authorization(roles.CreatePage), handler.CreatePageTxHandler(store))
	subrouter.PUT("/pages/:id", middleware.Authentication(server.TokenMaker), middleware.Authorization(roles.UpdatePage), handler.UpdatePagesTxHandler(store))
	subrouter.GET("/pages/:id", middleware.Authentication(server.TokenMaker), middleware.Authorization(roles.GetPage), handler.GetPageHandler(store))
	subrouter.GET("/pages", middleware.Authentication(server.TokenMaker), middleware.Authorization(roles.ListPages), handler.ListPagesHandler(store))
	subrouter.DELETE("/pages/:id", middleware.Authentication(server.TokenMaker), middleware.Authorization(roles.DeletePage), handler.DeletePageTxHandler(store))

	// Posts router
	subrouter.POST("/posts", middleware.Authentication(server.TokenMaker), middleware.Authorization(roles.CreatePost), handler.CreatePostTxHandler(store))
	subrouter.GET("/posts/:id", middleware.Authentication(server.TokenMaker), middleware.Authorization(roles.GetPage), handler.GetPostHandler(store))
	subrouter.GET("/posts", middleware.Authentication(server.TokenMaker), middleware.Authorization(roles.ListPosts), handler.ListPostsHandler(store))
	subrouter.PUT("/posts/:id", middleware.Authentication(server.TokenMaker), middleware.Authorization(roles.UpdatePost), handler.UpdatePostsTxHandler(store))
	subrouter.DELETE("/posts/:id", middleware.Authentication(server.TokenMaker), middleware.Authorization(roles.DeletePost), handler.DeletePostTxHandler(store))

	subrouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	server.Router = router
	return server, nil
}

func Start(server *common.Server, address string) error {
	return server.Router.Run(address)
}
