package service_jwt

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"diplomka/internal/model"

	"github.com/golang-jwt/jwt"
)

var sampleSecretKey = []byte("SecretYouShouldHide")

type jwtService struct {
	model.JWTService
}

func NewJWTService() model.JWTService {
	return &jwtService{}
}

func (j *jwtService) GenerateJWT(ctx context.Context, user model.User) (*model.Token, error) {
	claims := jwt.StandardClaims{
		Subject:   strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(sampleSecretKey)
	if err != nil {
		return nil, fmt.Errorf("error was ocured when creating token: %v", err)
	}

	return &model.Token{
		UserID:      user.ID,
		UserName:    user.Name,
		TokenString: signedToken,
	}, nil
}

func (j *jwtService) Verification(signedToken string) (int64, error) {
	token, err := jwt.ParseWithClaims(signedToken, &jwt.StandardClaims{}, keyFunc)
	if err != nil {
		return 0, fmt.Errorf("fwt parse: %v", err)
	}

	if !token.Valid {
		return 0, fmt.Errorf("token is not valid")
	}

	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		return 0, fmt.Errorf("invalid token")
	}

	userID, err := strconv.Atoi(claims.Subject)
	if err != nil {
		return 0, err
	}
	return int64(userID), nil
}

func keyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}

	return sampleSecretKey, nil
}
