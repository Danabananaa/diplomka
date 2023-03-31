package service

import (
	"context"
)

func (a *Auth) Refresh(ctx context.Context, token string) (string, error) {
	return "", nil
}
