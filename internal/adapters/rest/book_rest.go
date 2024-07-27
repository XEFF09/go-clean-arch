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

func (this_b *bookRestHandler) GetAllBooks(this_ctx *fiber.Ctx) error {
	
	books := this_b.bookUseCase.GetAllBooks()
	return this_ctx.Status(200).JSON(
		books,
	)
}

func (this_b *bookRestHandler) CreateBook(this_ctx *fiber.Ctx) error {
	var book entities.Book

	err := this_ctx.BodyParser(&book)
	if err != nil {
		return this_ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}
	createdBook, err := this_b.bookUseCase.CreateBook(book)
	
	if err != nil {
		switch err {
		case exceptions.DuplicateId:
			return this_ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
				"error": "Duplicate ID",
			})
		default:
			return this_ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Internal Server Error",
			})
		}
	}
	
	return this_ctx.Status(fiber.StatusCreated).JSON(
		createdBook,
	)
}