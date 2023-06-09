// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"go-fiber-postgres/controller"
	"go-fiber-postgres/repository"
	"go-fiber-postgres/service"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func NewAuth(db *gorm.DB) controller.AuthController {
	authRepository := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepository)
	authController := controller.NewAuthController(authService)
	return authController
}

func NewUser(db *gorm.DB) controller.UserController {
	userRepository := repository.NewUserRepository(db)
	authRepository := repository.NewAuthRepository(db)
	userService := service.NewUserService(userRepository, authRepository)
	userController := controller.NewUserController(userService)
	return userController
}


