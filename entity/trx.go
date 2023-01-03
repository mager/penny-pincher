package entity

import "time"

type Trx struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Timestamp time.Time `json:"ts"`
	PlaceID   string    `json:"place_id"`
	Amount    float64   `json:"amount"`
	BudgetID  int       `json:"budget_id"`
}
