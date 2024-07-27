package repositories

import "github.com/XEFF09/go-clean-arch/entities"

type BookRepository interface {
	GetAll() []entities.Book
	Create(book entities.Book) (*entities.Book, error)
}