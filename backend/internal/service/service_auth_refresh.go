package service

import (
	"context"
	"diplomka/internal/model"
	"fmt"
)

func (a *Auth) Refresh(ctx context.Context, token string) (string, error) {
	session := model.Session{
		Token: token,
	}
	fmt.Println(a.SessionRepo.IsSessionExist(ctx, session))
	return "", nil
}
