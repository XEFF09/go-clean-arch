package usecases

import (
	"github.com/XEFF09/go-clean-arch/entities"
	"github.com/XEFF09/go-clean-arch/repositories"
)

type BookUseCase interface {
	GetAllBooks() []entities.Book
	CreateBook(book entities.Book) (*entities.Book, error) 
}

type bookService struct {
	bookRepo repositories.BookRepository
}

func NewBookService(bookRepo repositories.BookRepository) BookUseCase {
	return &bookService{
		bookRepo: bookRepo,
	}
}

func (b *bookService) GetAllBooks() []entities.Book {
	return b.bookRepo.GetAll()
}
func (b *bookService) CreateBook(book entities.Book) (*entities.Book, error) {
	return b.bookRepo.Create(book)
}