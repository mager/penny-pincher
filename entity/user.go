package entity

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Phone    string `json:"phone"`
}
