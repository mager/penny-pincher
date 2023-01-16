package entity

type Budget struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
	Limit  int    `json:"limit"`
	Year   int    `json:"year"`
	Month  int    `json:"month"`
}
