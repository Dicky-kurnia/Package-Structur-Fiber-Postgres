package controller

import (
	"github.com/gofiber/fiber/v2"
)

type AuthController interface {
	Route(group fiber.Router)
	Login(c *fiber.Ctx) error
	Logout(c *fiber.Ctx) error
}
