package models

type OAuthLoginRequest struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}
