package common

import (
	"github.com/gin-gonic/gin"
	"github.com/reflection/frog-blossom-cms/config"
	db "github.com/reflection/frog-blossom-cms/db/sqlc"
	"github.com/reflection/frog-blossom-cms/token"
)

type Server struct {
	Store      db.Store
	TokenMaker token.Maker
	Config     config.Config
	Router     *gin.Engine
}
