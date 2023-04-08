package service

import (
	"context"
	"fmt"
	"os"

	"diplomka/internal/model"
)

type avatar struct {
	model.UserRepo
}

func NewAvatarService(ur model.UserRepo) *avatar {
	return &avatar{
		UserRepo: ur,
	}
}

func (a *avatar) AddUserImage(ctx context.Context, info model.UserImage) (*model.UserImage, error) {
	image, err := a.UserRepo.AddUserImage(ctx, info)
	if err != nil {
		return nil, fmt.Errorf("error was ocured from UserRepo AddUserImage: %v", err)
	}
	return image, nil
}

func (a *avatar) GetUserImageService(ctx context.Context, id int64) (*model.UserImage, error) {
	info, err := a.UserRepo.GetUserImage(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error was ocured from UserRepo GetUserImage: %v", err)
	}
	return info, nil
}

func (a *avatar) GetUserInfoService(ctx context.Context, id int) (*model.User, error) {
	user, err := a.UserRepo.GetUser(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error was ocured from UserRepo GetUser: %v", err)
	}
	return user, nil
}

func (a *avatar) DeleteImageService(ctx context.Context, id int64) error {
	ImgName, err := a.UserRepo.GetUserImage(ctx, id)
	if err != nil {
		return fmt.Errorf("error was ocured from UserRepo GetUserImage: %v", err)
	}
	path := "./static/images/" + ImgName.ImageName
	err = os.Remove(path)
	if err != nil {
		return err
	}
	err = a.UserRepo.DeleteImage(ctx, id)
	if err != nil {
		return fmt.Errorf("error was ocured from UserRepo DeleteImage: %v", err)
	}
	return nil
}
