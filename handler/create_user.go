package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/mager/penny-pincher/db"
	"github.com/mager/penny-pincher/entity"
)

func (h *Handler) createUser(w http.ResponseWriter, r *http.Request) {
	var (
		req  = entity.CreateUserReq{}
		resp = entity.CreateUserResp{}
		err  error
	)

	// Create a UUID for the user
	userID, _ := uuid.NewV4()
	// Get email from request context
	email := r.Context().Value("email").(string)

	// Decode request
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		handleServerError(err, w)
		return
	}

	q := db.CreateUserQuery(userID.String(), email, req.Phone, req.Country)
	h.Logger.Info(userID)
	h.Logger.Info(q)

	h.Logger.Infow("Running query", "handler", "createUser", "query", q)
	rows, err := h.Database.Query(h.Context, q)
	h.Logger.Info(rows)
	h.Logger.Info(err)
	if err != nil {
		handleServerError(err, w)
		return
	}

	h.Logger.Infow("User created", "handler", "createUser", "userid", userID, "email", email)

	resp.UserID = userID.String()
	resp.Email = email
	resp.Phone = req.Phone
	resp.Country = req.Country

	json.NewEncoder(w).Encode(resp)
}
