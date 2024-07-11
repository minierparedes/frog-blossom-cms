package common

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/reflection/frog-blossom-cms/config"
	db "github.com/reflection/frog-blossom-cms/db/sqlc"
	"github.com/reflection/frog-blossom-cms/token"
	"net/http"
)

func NewTestingServer(cfg config.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(cfg.TokenSystemmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		Store:      store,
		TokenMaker: tokenMaker,
		Config:     cfg,
	}
	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	subrouter := router.Group("api/v1")

	subrouter.POST("/users/login")

	subrouter.GET("/pages/:id", getUsersHandler(server.Store))
	subrouter.GET("/users/:id")
	subrouter.GET("/posts/:id")

	server.Router = router
}

type getUsersRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func getUsersHandler(store db.Store) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req getUsersRequest
		if err := ctx.ShouldBindUri(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		user, err := store.GetUsers(ctx, req.ID)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, err)
				return
			}

			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, user)
	}
}
