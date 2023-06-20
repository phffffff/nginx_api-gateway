package main

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type book struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	price       float32 `json:"price"`
	count       int     `json:"count"`
}

var books = []book{
	{
		Id:          1,
		Name:        "Sách 1",
		Description: "Mô tả sách 1",
		price:       34000,
		count:       1,
	},
	{
		Id:          2,
		Name:        "Sách 2",
		Description: "Mô tả sách 2",
		price:       56000,
		count:       2,
	},
	{
		Id:          3,
		Name:        "Sách 3",
		Description: "Mô tả sách 3",
		price:       45000,
		count:       3,
	},
}

func getBooks() []book {
	return books
}

func getBookById(id int) *book {
	for _, item := range books {
		if item.Id == id {
			return &item
		}
	}
	return nil
}

func handleGetBooks() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if len(books) > 0 {
			return c.Status(http.StatusOK).JSON(books)
		}
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"err:": errors.New("books invalid!")})
	}
}

func main() {
	app := fiber.New()
	app.Get("books", handleGetBooks())
	app.Listen(":8081")
}
