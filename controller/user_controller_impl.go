package controller

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go-fiber-postgres/helper"
	"go-fiber-postgres/model"
	"go-fiber-postgres/service"
)

type userController struct {
	service service.UserService
}

func NewUserController(service service.UserService) UserController {
	return &userController{service: service}
}

func (u userController) GetUserProfile(c *fiber.Ctx) error {
	var (
		ctx      = context.Background()
		username = c.Locals("currentUserName").(string)
	)
	response, err := u.service.GetUserProfile(ctx, username)
	helper.IsShouldPanic(err)

	return c.Status(200).JSON(model.Response{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (u userController) ChangePassword(c *fiber.Ctx) error {
	ctx := context.Background()
	request := new(model.ChangePasswordRequest)
	err := c.BodyParser(request)
	helper.IsShouldPanic(err)

	request.Username = c.Locals("currentUserName").(string)
	err = u.service.ChangePassword(ctx, *request)
	helper.IsShouldPanic(err)

	return c.Status(200).JSON(model.Response{
		Code:   200,
		Status: "OK",
	})
}

func (u userController) Route(group fiber.Router) {
	group.Get("", u.GetUserProfile)
	group.Put("/change-password", u.ChangePassword)
}
