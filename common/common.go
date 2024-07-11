package common

import (
	"github.com/gin-gonic/gin"
	"github.com/reflection/frog-blossom-cms/config"
	db "github.com/reflection/frog-blossom-cms/db/sqlc"
	"github.com/reflection/frog-blossom-cms/token"
)

// Server serves HTTP request for CMS
type Server struct {
	Store      db.Store
	TokenMaker token.Maker
	Config     config.Config
	Router     *gin.Engine
}

func ErrorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
