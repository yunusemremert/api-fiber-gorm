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
	Insert(book models.Book) (bool, error)
	GetAll() ([]models.Book, error)
	Get(id int16) (models.Book, error)
	Update(id int16) (bool, error)
	Delete(id int16) (bool, error)
}

func (bdb BookDB) Insert(book models.Book) (bool, error) {
	if result := bdb.DB.Create(book); result.Error != nil {
		log.Printf("bookRepository Insert error : %s", result.Error)

		return false, result.Error
	}

	return true, nil
}

func (bdb BookDB) GetAll() ([]models.Book, error) {
	var books []models.Book

	if result := bdb.DB.Find(books); result.Error != nil {
		log.Printf("bookRepository GetAll error : %s", result.Error)

		return nil, result.Error
	}

	return books, nil
}

func (bdb BookDB) Get() (models.Book, error) {
	var book models.Book

	return book, nil
}

func (bdb BookDB) Update() (bool, error) {
	return true, nil
}

func (bdb BookDB) Delete() (bool, error) {
	return true, nil
}
