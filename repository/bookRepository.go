package repository

import (
	"api-fiber-gorm/models"
	"gorm.io/gorm"
	"log"
)

type BookDB struct {
	DB *gorm.DB
}

type BookRepository interface {
	Insert(book models.Book) error
	GetAll() ([]models.Book, error)
	Get(id int16) (models.Book, error)
	Update(book models.Book, id int16) error
	Delete(id int16) error
}

func (bdb BookDB) Insert(book models.Book) error {
	if result := bdb.DB.Create(&book); result.Error != nil {
		log.Printf("bookRepository Insert error : %s", result.Error)

		return result.Error
	}

	return nil
}

func (bdb BookDB) GetAll() ([]models.Book, error) {
	var books []models.Book

	if result := bdb.DB.Find(&books); result.Error != nil {
		log.Printf("bookRepository GetAll error : %s", result.Error)

		return nil, result.Error
	}

	return books, nil
}

func (bdb BookDB) Get(id int16) (models.Book, error) {
	var book models.Book

	if result := bdb.DB.First(&book, id); result.Error != nil {
		log.Printf("bookRepository Get error : %s", result.Error)

		return models.Book{}, result.Error
	}

	return book, nil
}

func (bdb BookDB) Update(book models.Book, id int16) error {
	var oldBook models.Book

	if result := bdb.DB.First(&oldBook, id); result.Error != nil {
		log.Printf("bookRepository Update error : %s", result.Error)

		return result.Error
	}

	bdb.DB.Model(&oldBook).Updates(&book)

	return nil
}

func (bdb BookDB) Delete(id int16) error {
	var book models.Book

	if result := bdb.DB.First(&book, id); result.Error != nil {
		log.Printf("bookRepository Delete error : %s", result.Error)

		return result.Error
	}

	bdb.DB.Delete(&book, id)

	return nil
}
