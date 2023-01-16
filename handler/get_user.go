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
		resp   = entity.GetUserResp{}
		u      = entity.User{}
		userID = mux.Vars(r)["userID"]
		q      = db.GetUserQuery(userID)
		err    error
		rows   pgx.Rows
	)

	h.Logger.Infow("Running query", "handler", "getUser", "query", q)
	rows, err = h.Database.Query(h.Context, q)
	if err != nil {
		handleServerError(err, w)
		return
	}

	for rows.Next() {
		err = rows.Scan(&u.ID, &u.UserID, &u.Email, &u.Phone, &u.Locale)
		if err != nil {
			handleServerError(err, w)
			return
		}
	}
	if u.ID == 0 && resp.UserID == "" {
		w.WriteHeader(http.StatusNotFound)
	}

	// Adapt user
	resp.UserID = u.UserID
	resp.Email = u.Email
	resp.Phone = u.Phone
	resp.Locale = u.Locale

	json.NewEncoder(w).Encode(resp)
}
