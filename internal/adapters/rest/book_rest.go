package rest

import (
	"github.com/XEFF09/go-clean-arch/entities"
	"github.com/XEFF09/go-clean-arch/exceptions"
	usecases "github.com/XEFF09/go-clean-arch/usecases/book"
	"github.com/gofiber/fiber/v2"
)

type bookRestHandler struct {
	bookUseCase usecases.BookUseCase
}

func NewBookRestHandler(bookUseCase usecases.BookUseCase) bookRestHandler {
	return bookRestHandler{
		bookUseCase: bookUseCase,
	}
}

func (b *bookRestHandler) GetAllBooks(c *fiber.Ctx) error {
	
	books := b.bookUseCase.GetAllBooks()
	return c.Status(200).JSON(
		books,
	)
}

func (b *bookRestHandler) CreateBook(c *fiber.Ctx) error {
	var book entities.Book

	err := c.BodyParser(&book)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}
	createdBook, err := b.bookUseCase.CreateBook(book)
	
	if err != nil {
		switch err {
		case exceptions.DuplicateId:
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"error": "Duplicate ID",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Internal Server Error",
			})
		}
	}
	
	return c.Status(fiber.StatusCreated).JSON(
		createdBook,
	)
}