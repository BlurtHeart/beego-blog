package models

type UserLoginRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

type UserLoginResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
	Result   bool   `json:"result"`
}
