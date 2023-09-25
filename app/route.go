package app

import (
	"api-fiber-gorm/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(c *fiber.App) {
	route := c.Group("/api")

	v1 := route.Group("/v1")
	v1.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	bookController := controllers.BookService{}

	v1.Post("/book", bookController.CreateBook)
	v1.Get("/book", bookController.GetAllBook)
	v1.Get("/book/:id", bookController.GetBook)

	/*
		v2 := route.Group("/v2")
		v2.Get("/", func(c *fiber.Ctx) error {
			return c.SendString("Hello, World!")
		})
	*/
}
