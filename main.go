package main

import (
	r "api-fiber-gorm/app"
	"api-fiber-gorm/configs"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	configs.ConnectPostgreSQL()

	r.SetupRoutes(app)

	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
