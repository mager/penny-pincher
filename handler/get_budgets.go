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
	Foo     string          `json:"foo"`
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

	resp.Foo = "bar"
	h.Logger.Infow("Running query", "handler", "getBudgets", "query", q)
	rows, err = h.Database.Query(h.Context, q)
	h.Logger.Infow("DEBUG", "rows", rows, "resp", resp, "q", q, "handler", "getBudgets", "userID", userID)

	if err != nil {
		handleServerError(err, w)
		return
	}

	h.Logger.Info("No ERR")

	for rows.Next() {
		h.Logger.Info("Found something")
		var b entity.Budget
		err = rows.Scan(&b.ID, &b.UserID, &b.Name, &b.Limit, &b.Year)
		if err != nil {
			handleServerError(err, w)
			return
		}
		resp.Budgets = append(resp.Budgets, b)
	}

	json.NewEncoder(w).Encode(resp)
}
