package handler

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"go.uber.org/fx"

	"github.com/gorilla/mux"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v5"
	"github.com/mager/penny-pincher/config"
	"go.uber.org/zap"
)

type Handler struct {
	fx.In

	Context  context.Context
	Config   config.Config
	Database *pgx.Conn
	Logger   *zap.SugaredLogger
	Router   *mux.Router
}

// New creates a Handler struct
func New(h Handler) *Handler {
	h.registerRoutes()
	return &h
}

// RegisterRoutes registers all the routes for the route handler
func (h *Handler) registerRoutes() {
	// User
	h.Router.HandleFunc("/u/{id}", h.getUser).Methods("GET")
	h.Router.HandleFunc("/u", h.createUser).Methods("POST")

	// Budgets
	h.Router.HandleFunc("/b/{id}", h.getBudget).Methods("GET")
	h.Router.HandleFunc("/b/{id}/t", h.getBudgetTrxs).Methods("GET")
}

func handleServerError(err error, w http.ResponseWriter) {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		fmt.Println(pgErr.Message) // => syntax error at end of input
		w.Write([]byte(pgErr.Message))
		w.WriteHeader(http.StatusInternalServerError)
	}
}
