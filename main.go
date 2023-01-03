package main

import (
	c "context"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
	"github.com/mager/penny-pincher/config"
	"github.com/mager/penny-pincher/context"
	"github.com/mager/penny-pincher/db"
	"github.com/mager/penny-pincher/handler"
	"github.com/mager/penny-pincher/logger"
	"github.com/mager/penny-pincher/router"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.Provide(
			context.Options,
			config.Options,
			db.Options,
			logger.Options,
			router.Options,
		),
		fx.Invoke(Register),
	).Run()
}

func Register(
	lc fx.Lifecycle,
	ctx c.Context,
	cfg config.Config,
	db *pgx.Conn,
	log *zap.SugaredLogger,
	r *mux.Router,
) {
	p := handler.Handler{
		Context:  ctx,
		Config:   cfg,
		Database: db,
		Logger:   log,
		Router:   r,
	}

	handler.New(p)
}
