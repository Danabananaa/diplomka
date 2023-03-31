package service

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userid int64) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(1 * time.Hour)
	claims["userid"] = userid
	claims["authorized"] = true
	tokenString, err := token.SignedString(sampleSecretKey)
	if err != nil {
		return "", fmt.Errorf("Error was ocured when creating token: %v", err)
	}
	return tokenString, nil
}
