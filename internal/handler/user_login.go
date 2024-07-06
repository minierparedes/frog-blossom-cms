package handler

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/reflection/frog-blossom-cms/api"
	db "github.com/reflection/frog-blossom-cms/db/sqlc"
	"github.com/reflection/frog-blossom-cms/util"
	"net/http"
	"time"
)

type loginUsersRequest struct {
	ID       int64  `uri:"id" binding:"required,min=1"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

type loginUserResponse struct {
	AccessToken string       `json:"access_token"`
	User        userResponse `json:"user"`
}

type userResponse struct {
	ID          int64          `json:"id"`
	Username    string         `json:"username"`
	Email       string         `json:"email"`
	Role        string         `json:"role"`
	FirstName   string         `json:"first_name"`
	LastName    string         `json:"last_name"`
	UserUrl     sql.NullString `json:"user_url"`
	Description sql.NullString `json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		ID:          user.ID,
		Username:    user.Username,
		Email:       user.Email,
		Role:        user.Role,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		UserUrl:     user.UserUrl,
		Description: user.Description,
		CreatedAt:   user.CreatedAt,
	}
}

func LoginUser(server *api.Server, store db.Store) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req loginUsersRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}

		user, err := store.GetUsers(ctx, req.ID)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, errorResponse(err))
				return
			}
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		err = util.CheckPassword(req.Password, user.Password)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		accessToken, err := server.TokenMaker.CreateToken(user.Username, server.Config.AccessTokenDuration)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		resp := loginUserResponse{
			AccessToken: accessToken,
			User:        newUserResponse(user),
		}
		ctx.JSON(http.StatusOK, resp)
	}
}
