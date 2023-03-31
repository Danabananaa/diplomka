package model

type Token struct {
	UserID      int64  `json:"user_id"`
	TokenString string `json:"token"`
}
