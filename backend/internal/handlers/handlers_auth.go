package handlers

import "diplomka/internal/model"

type Auth struct {
	model.AuthSerivice
}

func NewAuthHandlers(a model.AuthSerivice) model.AuthHandlers {
	return &Auth{
		AuthSerivice: a,
	}
}
