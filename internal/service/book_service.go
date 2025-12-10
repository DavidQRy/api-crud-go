package service

import (
	"api-crud-go/internal/model"
	"api-crud-go/internal/store"
	"errors"
)

type Logger interface {
	Log(msg, error string)
}

type Service struct {
	store  store.Store
	logger Logger
}

func New(s store.Store) *Service {
	return &Service{
		store:  s,
		logger: nil,
	}
}
func (s *Service) GetAllBooks() ([]*model.Book, error) {
	s.logger.Log("Get the books", "")
	books, err := s.store.GetAll()
	if err != nil {
		s.logger.Log("The error is %v\n", err.Error())
		return nil, err
	}
	return books, nil
}

func (s *Service) GetBookById(id int) (*model.Book, error) {
	return s.store.GetByID(id)
}

func (s *Service) CreateBook(book model.Book) (*model.Book, error) {
	if book.Title == "" {
		return nil, errors.New("The title is required")
	}
	return s.store.Create(&book)
}

func (s *Service) UpdateBook(id int, book model.Book) (*model.Book, error) {
	if book.Title == "" {
		return nil, errors.New("The title is required")
	}
	return s.store.Update(id, &book)
}

func (s *Service) DeleteBook(id int) error {
	return s.store.Delete(id)
}
