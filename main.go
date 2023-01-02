package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mager/penny-pincher/logger"
	"github.com/mager/penny-pincher/router"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.Provide(
			logger.Options,
			router.Options,
		),
		fx.Invoke(Register),
	).Run()
}

func Register(lc fx.Lifecycle, log *zap.SugaredLogger, r *gin.Engine) {
	log.Info("Hello, World!")

	r.Run()
}
