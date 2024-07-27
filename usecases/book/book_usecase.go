package usecases

import (
	"github.com/XEFF09/go-clean-arch/entities"
	"github.com/XEFF09/go-clean-arch/repositories"
)

//Todo: module interface
type BookUseCase interface {
	GetAllBooks() []entities.Book
	CreateBook(book entities.Book) (*entities.Book, error) 
}

//Todo: module class that has repo
type bookService struct {
	bookRepo repositories.BookRepository
}

///Todo: module constructor
func NewBookService(bookRepo repositories.BookRepository) BookUseCase {
	//Todo: assign attr
	return &bookService{
		bookRepo: bookRepo,
	}
}

//Todo: module methods (thisb *bookService) as this tho
func (this_b *bookService) GetAllBooks() []entities.Book {
	return this_b.bookRepo.GetAll()
}
//Todo: module methods
func (this_b *bookService) CreateBook(book entities.Book) (*entities.Book, error) {
	return this_b.bookRepo.Create(book)
}