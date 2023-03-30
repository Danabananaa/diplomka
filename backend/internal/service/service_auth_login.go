package service

import (
	"context"
	"diplomka/internal/model"
	"fmt"
	"time"

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
	token, err := generateJWT(user.Name)
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

func generateJWT(name string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(1 * time.Hour)
	claims["authorized"] = true
	claims["user"] = name
	tokenString, err := token.SignedString(sampleSecretKey)
	if err != nil {
		return "", fmt.Errorf("Error was ocured when creating token: %v", err)
	}
	return tokenString, nil
}
