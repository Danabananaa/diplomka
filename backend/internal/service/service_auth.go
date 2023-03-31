package service

import "diplomka/internal/model"

type Auth struct {
	model.UserRepo
}

func NewAuthService(ur model.UserRepo) *Auth {
	return &Auth{
		UserRepo: ur,
	}
}
