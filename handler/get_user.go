package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
	"github.com/mager/penny-pincher/db"
	"github.com/mager/penny-pincher/entity"
)

func (h *Handler) getUser(w http.ResponseWriter, r *http.Request) {
	var (
		resp    = entity.GetUserResp{}
		auth0ID = mux.Vars(r)["id"]
		q       = db.GetUserQuery(auth0ID)
		err     error
		rows    pgx.Rows
	)

	h.Logger.Infow("Running query", "handler", "getUser", "query", q)
	rows, err = h.Database.Query(h.Context, q)
	if err != nil {
		handleServerError(err, w)
		return
	}

	for rows.Next() {
		err = rows.Scan(&resp.ID, &resp.Auth0ID, &resp.Username, &resp.Email)
		if err != nil {
			handleServerError(err, w)
			return
		}
	}
	if resp.ID == 0 && resp.Auth0ID == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(resp)
}
