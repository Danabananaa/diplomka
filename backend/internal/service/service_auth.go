package service

import "diplomka/internal/model"

type Auth struct {
	model.UserRepo
	model.SessionRepo
}

func NewAuthService(ur model.UserRepo, sr model.SessionRepo) *Auth {
	return &Auth{
		UserRepo:    ur,
		SessionRepo: sr,
	}
}
