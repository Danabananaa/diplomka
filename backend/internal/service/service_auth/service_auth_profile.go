package service_auth

import (
	"context"

	"diplomka/internal/model"
)

func (a *auth) Profile(ctx context.Context, userId int64) (model.Profile, error) {
	profile := model.Profile{}

	user, err := a.UserRepo.GetUser(ctx, int(userId))
	if err != nil {
		return profile, err
	}

	profile.ID = user.ID
	profile.Name = user.Name
	profile.Surname = user.Surname
	profile.Email = user.Email

	return profile, nil
}
