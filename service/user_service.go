package service

import (
	"context"
	"go-fiber-postgres/model"
)

type UserService interface {
	GetUserProfile(ctx context.Context, username string) (*model.GetUserProfileResponse, error)
	ChangePassword(ctx context.Context, request model.ChangePasswordRequest) error
}
