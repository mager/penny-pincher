package handler

import (
	"encoding/json"
	"net/http"

	"github.com/mager/penny-pincher/db"
	"github.com/mager/penny-pincher/entity"
)

func (h *Handler) createUser(w http.ResponseWriter, r *http.Request) {
	var (
		req  = entity.CreateUserReq{}
		resp = entity.CreateUserResp{}
		err  error
	)

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		handleServerError(err, w)
		return
	}
	q := db.CreateUserQuery(req)

	h.Logger.Infow("Running query", "handler", "createUser", "query", q)
	_, err = h.Database.Query(h.Context, q)
	if err != nil {
		handleServerError(err, w)
		return
	}

	resp.UserID = req.UserID
	resp.Phone = req.Phone
	resp.Email = req.Email
	resp.Locale = req.Locale

	json.NewEncoder(w).Encode(resp)
}
