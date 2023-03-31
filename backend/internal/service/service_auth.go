package service

import "diplomka/internal/model"

type auth struct {
	model.UserRepo
	model.JWTService
}

func NewAuthService(ur model.UserRepo, js model.JWTService) *auth {
	return &auth{
		UserRepo:   ur,
		JWTService: js,
	}
}
