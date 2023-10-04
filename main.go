package main

import (
	r "api-fiber-gorm/app"
	"api-fiber-gorm/configs"
	"api-fiber-gorm/controllers"
	"api-fiber-gorm/repository"
	"api-fiber-gorm/services"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	dbClient := configs.ConnectPostgreSQL()

	bookRepoDB := repository.NewBookRepositoryDB(dbClient)
	bookService := services.NewBookService(bookRepoDB)
	bookController := controllers.NewBookController(bookService)

	r.SetupBookRoutes(app, bookController)
	// other routes

	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
