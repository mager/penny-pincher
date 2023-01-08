package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
	"github.com/mager/penny-pincher/db"
	"github.com/mager/penny-pincher/entity"
)

func (h *Handler) getBudgetTrxs(w http.ResponseWriter, r *http.Request) {
	var (
		id   = mux.Vars(r)["id"]
		q    = db.GetBudgetTransactionsQuery(id)
		err  error
		rows pgx.Rows
	)

	h.Logger.Infow("Running query", "handler", "getBudgetTrxs", "query", q)
	rows, err = h.Database.Query(h.Context, q)
	if err != nil {
		handleServerError(err, w)
		return
	}

	// Collect rows
	resp, err := pgx.CollectRows(
		rows,
		func(row pgx.CollectableRow) (entity.Trx, error) {
			var t entity.Trx
			err = rows.Scan(&t.ID, &t.UserID, &t.Timestamp, &t.PlaceID, &t.Amount, &t.BudgetID)
			if err != nil {
				handleServerError(err, w)
			}
			return t, nil
		},
	)

	json.NewEncoder(w).Encode(resp)
}
