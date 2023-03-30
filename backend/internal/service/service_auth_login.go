package service

import (
	"context"
	"diplomka/internal/model"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

var sampleSecretKey = []byte("SecretYouShouldHide")

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (a *Auth) LogIn(ctx context.Context, email, password string) (string, error) {
	user, err := a.UserRepo.GetUserforAuth(ctx, email, password)
	if err != nil {
		return "", fmt.Errorf("Error was ocured from UserRepo GetUserforAuth: %v", err)
	}
	token, err := GenerateJWT()
	if err != nil {
		return "", err
	}

	session := model.Session{
		Token:  token,
		UserID: user.ID,
	}
	err = a.SessionRepo.AddSession(ctx, session)
	if err != nil {
		return "", fmt.Errorf("Error was ocured from SesionRepo AddSession: %v", err)
	}

	return token, nil
}
