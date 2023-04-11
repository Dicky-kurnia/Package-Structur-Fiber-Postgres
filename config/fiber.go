package config

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"go-fiber-postgres/exception"
)

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
		ErrorHandler: exception.ErrorHandler,
	}
}
