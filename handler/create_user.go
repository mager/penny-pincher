package handler

import (
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/mager/penny-pincher/db"
	"github.com/mager/penny-pincher/entity"
)

func (h *Handler) createUser(w http.ResponseWriter, r *http.Request) {
	var (
		req  = entity.CreateUserReq{}
		resp = entity.CreateUserResp{}
		u    = entity.User{}
		err  error
		rows pgx.Rows
	)

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		handleServerError(err, w)
		return
	}
	q := db.CreateUserQuery(req)

	h.Logger.Info(q)
	h.Logger.Info(req)

	h.Logger.Infow("Running query", "handler", "createUser", "query", q)
	rows, err = h.Database.Query(h.Context, q)
	if err != nil {
		handleServerError(err, w)
		return
	}

	for rows.Next() {
		err = rows.Scan(&u.ID, &u.Auth0ID, &u.Email, &u.Phone, &u.Locale)
		if err != nil {
			handleServerError(err, w)
			return
		}
	}
	if u.ID == 0 && resp.Auth0ID == "" {
		w.WriteHeader(http.StatusNotFound)
	}

	json.NewEncoder(w).Encode(resp)
}
