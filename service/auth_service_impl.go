package service

import (
	"context"
	"errors"
	"go-fiber-postgres/helper"
	"go-fiber-postgres/model"
	"go-fiber-postgres/repository"
	"go-fiber-postgres/validation"

	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	repository repository.AuthRepository
}

func NewAuthService(repository repository.AuthRepository) AuthService {
	return &authService{repository: repository}
}

func (service *authService) Login(ctx context.Context, request *model.LoginRequest) (response *model.LoginResponse, salesData *model.Sales, err error) {
	cek, err := validation.Validate(request)

	if cek {
		return response, salesData, err
	}

	salesData, err = service.repository.FetchSales(ctx, request.Username)
	if err != nil {
		return response, salesData, errors.New(model.USERNAME_OR_PASSWORD_INVALID)
	}

	err = bcrypt.CompareHashAndPassword([]byte(salesData.Password), []byte(request.Password))

	if err != nil {
		return &model.LoginResponse{}, salesData, errors.New(model.USERNAME_OR_PASSWORD_INVALID)
	}

	jwtPayload := model.JwtPayload{
		SalesId:  salesData.Id,
		Username: salesData.Username,
	}
	//generate token
	ts := helper.CreateToken(jwtPayload)

	//save metadata to redis
	helper.CreateAuth(jwtPayload, ts)

	response = &model.LoginResponse{
		AccessToken: ts.AccessToken,
	}

	if err != nil {
		return response, salesData, err
	}

	return response, salesData, nil
}

func (service *authService) Logout(token string) (err error) {
	helper.DelRedis(token)

	return nil
}
