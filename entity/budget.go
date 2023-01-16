package entity

type Budget struct {
	ID     int32  `json:"id"`
	UserID string `json:"userid"`
	Name   string `json:"name"`
	Limit  int32  `json:"limit"`
	Year   int32  `json:"year"`
	Month  int32  `json:"month"`
}
