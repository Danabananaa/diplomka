package service

import (
	"context"
	"fmt"

	"diplomka/internal/model"
)

func (a *auth) SignUp(ctx context.Context, user model.User) error {
	_, err := a.UserRepo.AddUser(ctx, user)
	if err != nil {
		return fmt.Errorf("user repo AddUser: %v", err)
	}
	return nil
}
