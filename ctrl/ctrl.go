package ctrl

import (
	"context"

	"go.uber.org/fx"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/mager/penny-pincher/config"
	"go.uber.org/zap"
)

type Controller struct {
	fx.In

	Context  context.Context
	Config   config.Config
	Database *pgx.Conn
	Logger   *zap.SugaredLogger
	Router   *gin.Engine
}

func New(c Controller) *Controller {
	return &c
}
