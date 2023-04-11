//go:build wireinject
// +build wireinject

package wire

import (
	"api-dev/controller"
	"api-dev/repository"
	"api-dev/service"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func NewAuth(db *gorm.DB) controller.AuthController {
	wire.Build(repository.NewAuthRepository, service.NewAuthService, controller.NewAuthController)
	return nil
}

func NewUser(db *gorm.DB) controller.UserController {
	wire.Build(repository.NewUserRepository, repository.NewAuthRepository, service.NewUserService, controller.NewUserController)
	return nil
}
