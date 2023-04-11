package service

import (
	"context"
	"go-fiber-postgres/model"
)

type AuthService interface {
	Login(ctx context.Context, request *model.LoginRequest) (response *model.LoginResponse, salesData *model.Sales, err error)
	Logout(token string) (err error)
}
