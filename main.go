package main

import (
	"github.com/XEFF09/go-clean-arch/internal/adapters/memory"
	"github.com/XEFF09/go-clean-arch/internal/adapters/rest"
	usecases "github.com/XEFF09/go-clean-arch/usecases/book"
	"github.com/gofiber/fiber/v2"
)

func main() {

	bookRepo := memory.NewBookMemory()
	bookService := usecases.NewBookService(bookRepo)
	bookHandler := rest.NewBookRestHandler(bookService)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/books", bookHandler.GetAllBooks)
	app.Post("/books", bookHandler.CreateBook)

	app.Listen(":3000")
}