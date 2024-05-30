package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/reflection/frog_blossom_db/db/sqlc"
)

// CreateUsers handler

type createUsersRequest struct {
	Username    string `json:"username" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Role        string `json:"role" binding:"required"`
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	UserUrl     string `json:"user_url" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func (server *Server) createUsers(ctx *gin.Context) {
	var req createUsersRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := db.CreateUsersParams{
		Username:    req.Username,
		Email:       req.Email,
		Password:    req.Password,
		Role:        req.Role,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		UserUrl:     sql.NullString{String: req.UserUrl, Valid: true},
		Description: sql.NullString{String: req.Description, Valid: true},
	}

	user, err := server.store.CreateUsers(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}
