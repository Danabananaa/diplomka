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

	token, err := a.JWTService.GenerateJWT(ctx, *user)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (a *auth) GetUserInfoService(ctx context.Context, id int) (*model.User, error) {
	user, err := a.UserRepo.GetUser(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error was ocured from UserRepo GetUser: %v", err)
	}
	return user, nil
}

func (a *auth) AddUserImage(ctx context.Context, info model.UserImage) (*model.UserImage, error) {
	image, err := a.UserRepo.AddUserImage(ctx, info)
	if err != nil {
		return nil, fmt.Errorf("error was ocured from UserRepo AddUserImage: %v", err)
	}
	return image, nil
}

func (a *auth) GetUserImageService(ctx context.Context, id int) (*model.UserImage, error) {
	info, err := a.UserRepo.GetUserImage(ctx, id)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("error was ocured from UserRepo GetUserImage: %v", err)
	}
	return info, nil
}
