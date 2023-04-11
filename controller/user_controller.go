package controller

import "github.com/gofiber/fiber/v2"

type UserController interface {
	GetUserProfile(c *fiber.Ctx) error
	Route(group fiber.Router)
}
