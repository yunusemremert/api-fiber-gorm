package services

import (
	"api-fiber-gorm/models"
	"api-fiber-gorm/repository"
)

type BookRepository struct {
	Repository repository.BookRepository
}

type BookService interface {
	BookInsert(book models.Book) error
	BookGetAll() ([]models.Book, error)
	BookGet(id int16) (models.Book, error)
	BookUpdate(book models.Book, id int16) error
	BookDelete(id int16) error
}

func (br BookRepository) BookInsert(book models.Book) error {
	err := br.Repository.Insert(book)
	if err != nil {
		return err
	}

	return nil
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

func (br BookRepository) BookUpdate(book models.Book, id int16) error {
	err := br.Repository.Update(book, id)
	if err != nil {
		return err
	}

	return nil
}

func (br BookRepository) BookDelete(id int16) error {
	err := br.Repository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func NewBookService(repository repository.BookRepository) *BookRepository {
	return &BookRepository{repository}
}
