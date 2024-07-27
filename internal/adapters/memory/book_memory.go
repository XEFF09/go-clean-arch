package memory

import (
	"github.com/XEFF09/go-clean-arch/entities"
	"github.com/XEFF09/go-clean-arch/exceptions"
	"github.com/XEFF09/go-clean-arch/repositories"
)

type BookMemory struct {
	books []entities.Book
}

// Create implements repositories.BookRepository.
func (b *BookMemory) Create(book entities.Book) (*entities.Book, error) {
	
	for _, b := range b.books {
		if b.ID == book.ID {
			return nil, exceptions.DuplicateId
		}
	}
	
	b.books = append(b.books, book)
	return &book, nil
}

func (b *BookMemory) GetAll() []entities.Book {
	return b.books
}

func NewBookMemory() repositories.BookRepository {
	return &BookMemory{
		books: []entities.Book{},
	}
}
