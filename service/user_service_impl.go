package service

import (
	"context"
	"errors"
	"go-fiber-postgres/exception"
	"go-fiber-postgres/model"
	"go-fiber-postgres/repository"
	"go-fiber-postgres/validation"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type userService struct {
	repository     repository.UserRepository
	AuthRepository repository.AuthRepository
}

func NewUserService(repository repository.UserRepository, authRepository repository.AuthRepository) UserService {
	return &userService{repository: repository, AuthRepository: authRepository}
}

func (u userService) GetUserProfile(ctx context.Context, username string) (*model.GetUserProfileResponse, error) {
	return u.repository.GetUserProfile(ctx, username)
}

func (u userService) ChangePassword(ctx context.Context, request model.ChangePasswordRequest) error {
	cek, err := validation.Validate(request)

	if cek {
		return err
	}

	if strings.Compare(request.OldPassword, request.NewPassword) == 0 {
		return errors.New(model.OLD_PASSWORD_CANNOT_BE_MATCH)
	}

	salesData, err := u.AuthRepository.FetchSales(ctx, request.Username)
	if err != nil {
		return exception.NewValidationError("old_password", model.NOT_VALID_ERR_TYPE)
	}

	err = bcrypt.CompareHashAndPassword([]byte(salesData.Password), []byte(request.OldPassword))
	if err != nil {
		return exception.NewValidationError("old_password", model.NOT_VALID_ERR_TYPE)

	}

	if strings.Compare(request.NewPassword, request.ConfirmNewPassword) != 0 {
		return exception.CONFIRM_PASSWORD_NOT_MATCH
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(request.NewPassword), 10)
	if err != nil {
		return exception.NewValidationError("new_password", model.NOT_VALID_ERR_TYPE)
	}

	username, err := u.repository.ChangePassword(ctx, request.Username, string(pass))
	if username == "" || err != nil {
		return exception.NewValidationError("old_password", model.NOT_VALID_ERR_TYPE)
	}

	return nil
}
