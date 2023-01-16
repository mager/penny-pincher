package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
	"github.com/mager/penny-pincher/db"
	"github.com/mager/penny-pincher/entity"
)

type GetBudgetsResp struct {
	Budgets []entity.Budget `json:"budgets"`
}

func (h *Handler) getBudgets(w http.ResponseWriter, r *http.Request) {
	var (
		budgets = make([]entity.Budget, 0)
		resp    = GetBudgetsResp{
			Budgets: budgets,
		}
		userID = mux.Vars(r)["userID"]
		q      = db.GetBudgetsQuery(userID)
		err    error
		rows   pgx.Rows
	)

	h.Logger.Infow("Running query", "handler", "getBudgets", "query", q)
	rows, err = h.Database.Query(h.Context, q)

	if err != nil {
		handleServerError(err, w)
		return
	}

	for rows.Next() {
		var b entity.Budget
		err = rows.Scan(&b.ID, &b.UserID, &b.Name, &b.Limit, &b.Year, &b.Month)
		if err != nil {
			handleServerError(err, w)
			return
		}
		resp.Budgets = append(resp.Budgets, b)
	}

	json.NewEncoder(w).Encode(resp)
}
