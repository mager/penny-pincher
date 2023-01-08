package entity

type GetUserResp struct {
	ID       int    `json:"id"`
	Auth0ID  string `json:"auth0_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
