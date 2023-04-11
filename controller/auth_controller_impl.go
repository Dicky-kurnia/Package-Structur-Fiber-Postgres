package controller

import (
	"context"
	"go-fiber-postgres/helper"
	"go-fiber-postgres/model"
	"go-fiber-postgres/service"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type authController struct {
	service service.AuthService
}

func NewAuthController(service service.AuthService) AuthController {
	return &authController{service}
}

func (controller *authController) Route(group fiber.Router) {
	group.Post("", controller.Login)
	group.Post("/logout", controller.Logout)
}

func (controller *authController) Login(c *fiber.Ctx) error {
	request := new(model.LoginRequest)

	if err := c.BodyParser(request); err != nil {
		helper.IsShouldPanic(err)
	}

	response, _, err := controller.service.Login(context.Background(), request)
	helper.IsShouldPanic(err)

	return c.Status(200).JSON(model.Response{
		Code:   200,
		Status: "OK",
		Data:   response,
	})

}

func (controller *authController) Logout(c *fiber.Ctx) error {
	tokenSlice := strings.Split(c.Get("Authorization"), "Bearer ")

	var tokenString string
	if len(tokenSlice) == 2 {
		tokenString = tokenSlice[1]
	}

	err := controller.service.Logout(tokenString)

	if err != nil {
		helper.IsShouldPanic(err)
	}

	return c.Status(200).JSON(model.Response{
		Code:   200,
		Status: "OK",
	})
}
