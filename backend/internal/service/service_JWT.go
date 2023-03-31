package service

import (
	"context"
	"fmt"
	"time"

	"diplomka/internal/model"

	"github.com/golang-jwt/jwt/v5"
)

var sampleSecretKey = []byte("SecretYouShouldHide")

type jwtService struct {
	model.JWTService
}

func NewJWTService() *jwtService {
	return &jwtService{}
}

func (j *jwtService) GenerateJWT(ctx context.Context, userID int64) (*model.Token, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"iat": time.Now(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(sampleSecretKey)
	if err != nil {
		return nil, fmt.Errorf("error was ocured when creating token: %v", err)
	}

	return &model.Token{
		UserID:      userID,
		TokenString: signedToken,
	}, nil
}

func (j *jwtService) Verification(signedToken string) error {
	token, err := jwt.Parse(signedToken, keyFunc)
	if err != nil {
		return fmt.Errorf("fwt parse: %v", err)
	}

	if !token.Valid {
		return fmt.Errorf("token is not valid")
	}

	return nil
}

func keyFunc(token *jwt.Token) (interface{}, error) {
	// fmt.Printf("token.Claims: %v\n", token.Claims)
	// fmt.Printf("token.Header: %v\n", token.Header)
	// fmt.Printf("token.Method: %v\n", token.Method)
	// fmt.Printf("token.Raw: %v\n", token.Raw)
	// fmt.Printf("token.Signature: %v\n", token.Signature)
	// fmt.Printf("token.Valid: %v\n", token.Valid)
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}

	// проверка токена будет писаться здесь
	// проверка на время и другие значения
	// посмотри какие значения находятся в токене

	return sampleSecretKey, nil
}
