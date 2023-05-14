package entity

type UserFields struct {
	UserID  string `json:"userid"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Country string `json:"country"`
}

type User struct {
	ID int `json:"id"`
	UserFields
}

type GetUserReq struct {
	Email string `json:"email"`
}

type GetUserResp struct {
	UserFields
}

type CreateUserReq struct {
	Phone   string `json:"phone"`
	Country string `json:"country"`
}

type CreateUserResp struct {
	UserFields
}
