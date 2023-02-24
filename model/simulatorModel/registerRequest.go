package simModels

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Pwd      string `json:"pwd"`
}
