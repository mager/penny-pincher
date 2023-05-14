package handler

import (
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/mager/penny-pincher/db"
	"github.com/mager/penny-pincher/entity"
)

func (h *Handler) getUser(w http.ResponseWriter, r *http.Request) {
	var (
		resp = entity.GetUserResp{}
		u    = entity.User{}
		err  error
		rows pgx.Rows
	)

	email := r.Context().Value("email").(string)
	q := db.GetUserQuery(email)

	h.Logger.Infow("Running query", "handler", "getUser", "query", q)
	rows, err = h.Database.Query(h.Context, q)
	if err != nil {
		handleServerError(err, w)
		return
	}

	for rows.Next() {
		err = rows.Scan(&u.ID, &u.UserID, &u.Email, &u.Phone, &u.Country)
		if err != nil {
			handleServerError(err, w)
			return
		}
	}

	if u.ID == int(0) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Adapt user
	resp.UserID = u.UserID
	resp.Email = u.Email

	json.NewEncoder(w).Encode(resp)
}
