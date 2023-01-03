package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/mager/penny-pincher/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func ProvideDB(
	lc fx.Lifecycle,
	ctx context.Context,
	cfg config.Config,
	log *zap.SugaredLogger,
) *pgx.Conn {
	var (
		userName   = "mager"
		dbName     = "mager/penny-pincher"
		connectUrl = fmt.Sprintf("postgres://%s:%s@db.bit.io:5432/%s", userName, cfg.BitIOAPIKey, dbName)
	)

	conn, err := pgx.Connect(ctx, connectUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return conn
}

var Options = ProvideDB

func GetBudgetByIDQuery(id string) string {
	return fmt.Sprintf("SELECT * FROM budgets WHERE id = %s", id)
}
