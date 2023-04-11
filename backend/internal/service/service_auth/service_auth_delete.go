package service_auth

import (
	"context"
	"fmt"
)

func (a *auth) Delete(ctx context.Context, userID int64) error {
	err := a.UserRepo.DeleteUser(ctx, userID)
	if err != nil {
		return fmt.Errorf("error deleting user %d: %v", userID, err)
	}
	return nil
}
