package entity

type Budget struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	UserID int    `json:"user_id"`
	Limit  int    `json:"limit"`
}
