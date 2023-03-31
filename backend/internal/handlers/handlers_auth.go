package handlers

import "diplomka/internal/model"

type auth struct {
	model.AuthSerivice
}

func NewAuthHandlers(a model.AuthSerivice) *auth {
	return &auth{
		AuthSerivice: a,
	}
}
