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

	bc := controllers.BookService{}

	v1.Post("/book", bc.CreateBook)
	v1.Get("/book", bc.GetAllBook)
	v1.Get("/book/:id", bc.GetBook)
	v1.Put("/book/:id", bc.UpdateBook)
	v1.Delete("/book/:id", bc.DeleteBook)

	/*
		v2 := route.Group("/v2")
		v2.Get("/", func(c *fiber.Ctx) error {
			return c.SendString("Hello, World!")
		})
	*/
}
