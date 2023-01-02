package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/mager/penny-pincher/config"
	"go.uber.org/zap"
)

func ProvideDB(cfg config.Config, log *zap.SugaredLogger) *pgx.Conn {
	log.Info("Connecting to database...")
	userName := "mager"
	dbName := "mager/penny-pincher"
	connectUrl := fmt.Sprintf("postgres://%s:%s@db.bit.io:5432/%s", userName, cfg.BitIOAPIKey, dbName)
	conn, err := pgx.Connect(context.Background(), connectUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	return conn
}

var Options = ProvideDB
