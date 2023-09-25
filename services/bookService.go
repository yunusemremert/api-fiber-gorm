package services

import (
	"api-fiber-gorm/models"
	"api-fiber-gorm/repository"
)

type BookRepository struct {
	Repository repository.BookRepository
}

type BookService interface {
	BookInsert(book models.Book) (bool, error)
	BookGetAll() ([]models.Book, error)
	BookGet(id int16) (models.Book, error)
	BookUpdate(id int16) (bool, error)
	BookDelete(id int16) (bool, error)
}

func (br BookRepository) BookInsert(book models.Book) (bool, error) {
	result, err := br.Repository.Insert(book)
	if result == false || err != nil {
		return false, err
	}

	return true, nil
}

func (br BookRepository) BookGetAll() ([]models.Book, error) {
	result, err := br.Repository.GetAll()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (br BookRepository) BookGet(id int16) (models.Book, error) {
	result, err := br.Repository.Get(id)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (br BookRepository) BookUpdate(id int16) (bool, error) {
	return true, nil
}

func (br BookRepository) BookDelete(id int16) (bool, error) {
	return true, nil
}
