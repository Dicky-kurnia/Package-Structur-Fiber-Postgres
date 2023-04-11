package repository

import (
	"context"
	"go-fiber-postgres/model"
)

type UserRepository interface {
	GetSalesIdByUsername(ctx context.Context, username string) (uint, error)
	GetUserProfile(ctx context.Context, username string) (*model.GetUserProfileResponse, error)
	ChangePassword(ctx context.Context, username, password string) (string, error)
}
