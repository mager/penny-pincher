package entity

import "time"

type Trx struct {
	ID        int       `json:"id"`
	BudgetID  int       `json:"budget_id"`
	Amount    int       `json:"amount"`
	Timestamp time.Time `json:"ts"`
	PlaceID   int       `json:"place_id"`
	UserID    int       `json:"user_id"`
}
