package controllers

import (
	"api-fiber-gorm/models"
	"api-fiber-gorm/services"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type BookService struct {
	Service services.BookService
}

func (bs BookService) CreateBook(c *fiber.Ctx) error {
	var book models.Book

	if err := c.BodyParser(&book); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid Json.",
			"message": err.Error(),
		})
	}

	result, err := bs.Service.BookInsert(book)
	if err != nil {
		return c.Status(http.StatusBadGateway).JSON(fiber.Map{
			"error":   "Service unavailable",
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Book created successfully.",
		"data":    result,
	})
}

func (bs BookService) GetAllBook(c *fiber.Ctx) error {
	result, err := bs.Service.BookGetAll()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(result)
}

func (bs BookService) GetBook(c *fiber.Ctx) error {
	return nil
}
