package main

import (
	"go-fiber-postgres/config"
	"go-fiber-postgres/helper"
	"go-fiber-postgres/middleware"
	"go-fiber-postgres/wire"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/joho/godotenv"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	err := godotenv.Load()
	if err != nil {
		helper.IsShouldPanic(err)
	}
	loc, _ := time.LoadLocation("Asia/Jakarta")
	time.Local = loc
}

func main() {
	// Declare new Database Object
	db := config.NewFMCGPostgresDB()

	// User Fields
	authController := wire.NewAuth(db)

	userController := wire.NewUser(db)

	// Initialize Fiber
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())
	app.Static("/public", "./storage/public")

	v1 := app.Group("/")

	auth := v1.Group("/auth")
	user := v1.Group("/user", middleware.CheckToken)

	// User Controller Routing
	authController.Route(auth)

	userController.Route(user)

	// Fiber Start
	err := app.Listen(":" + os.Getenv("PORT"))
	helper.IsShouldPanic(err)
}
