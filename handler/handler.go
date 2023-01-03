package handler

import (
	"context"
	"net/http"

	"go.uber.org/fx"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
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
	// Update collections
	h.Router.HandleFunc("/b/{id}", h.getBudget).Methods("GET")
}

func handleServerError(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}
