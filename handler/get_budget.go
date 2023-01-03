package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
	"github.com/mager/penny-pincher/db"
	"github.com/mager/penny-pincher/entity"
)

func (h *Handler) getBudget(w http.ResponseWriter, r *http.Request) {
	var (
		resp = entity.Budget{}
		id   = mux.Vars(r)["id"]
		q    = db.GetBudgetByIDQuery(id)
		err  error
		rows pgx.Rows
	)

	rows, err = h.Database.Query(h.Context, q)
	if err != nil {
		handleServerError(err, w)
		return
	}

	for rows.Next() {
		err = rows.Scan(&resp.ID, &resp.UserID, &resp.Name, &resp.Limit, &resp.Year)
		if err != nil {
			handleServerError(err, w)
			return
		}
	}

	json.NewEncoder(w).Encode(resp)
}
