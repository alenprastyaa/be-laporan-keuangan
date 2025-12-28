package controllers

import (
	"laporan-keuangan/models"
	"laporan-keuangan/utils"

	"github.com/gofiber/fiber/v2"
)

func GetBooks(c *fiber.Ctx) error {
	books, err := models.GetAllBook()
	if err != nil {
		return utils.ResponseError(c, fiber.StatusInternalServerError, "Gagal Mengambil data buku")
	}
	if books == nil {
		books = []models.Book{}
	}
	return utils.ResponseSuccess(c, books, "Data buku berhasl diambil")
}

func CreateBook(c *fiber.Ctx) error {
	var book models.Book
	if err := c.BodyParser(&book); err != nil {
		return utils.ResponseError(c, fiber.StatusBadRequest, "input tidak valid")
	}
	userId := c.Locals("user_id").(float64)
	book.UserID = int(userId)
	if err := models.CreateBook(&book); err != nil {
		return utils.ResponseError(c, fiber.StatusInternalServerError, "gagal Menambahkan buku")
	}
	return utils.ResponseSuccess(c, book, "Buku berhasil ditambahkan")
}

func GetMyBook(c *fiber.Ctx) error {
	userVal := c.Locals("user_id")
	if userVal == nil {
		return utils.ResponseError(c, fiber.StatusUnauthorized, "User ID tidak valid")
	}
	userID := int(userVal.(float64))
	books, err := models.GetBooksByUserID(userID)
	if err != nil {
		return utils.ResponseError(c, fiber.StatusInternalServerError, "Gagal mengambil data buku")
	}
	if books == nil {
		books = []models.Book{}
	}
	return utils.ResponseSuccess(c, books, "Data buku berhasil diambil")
}
