package entity

type UserFields struct {
	Auth0  string `json:"auth0"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`
	Locale string `json:"locale"`
}

type User struct {
	ID int `json:"id"`
	UserFields
}

type GetUserResp struct {
	UserFields
}

type CreateUserReq struct {
	UserFields
}

type CreateUserResp struct {
	UserFields
}
