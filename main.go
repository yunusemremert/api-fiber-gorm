package main

import (
	r "api-fiber-gorm/app"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	r.SetupRoutes(app)

	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
