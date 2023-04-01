package service

import (
	"context"
	"diplomka/internal/model"
	"fmt"
)

func (a *auth) LogIn(ctx context.Context, auth model.Authentication) (*model.Token, error) {
	user, err := a.UserRepo.GetUserforAuth(ctx, auth)
	if err != nil {
		return nil, fmt.Errorf("error was ocured from UserRepo GetUserforAuth: %v", err)
	}

	token, err := a.JWTService.GenerateJWT(ctx, user.ID, user.Name)
	if err != nil {
		return nil, err
	}

	return token, nil
}
