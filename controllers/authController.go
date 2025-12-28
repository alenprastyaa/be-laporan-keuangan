package controllers

import (
	"laporan-keuangan/models"
	"laporan-keuangan/utils"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var input models.User
	if err := c.BodyParser(&input); err != nil {
		return utils.ResponseError(c, fiber.StatusBadRequest, "Invalid input data")
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	input.Password = string(hashedPassword)

	if _, err := models.CreateUser(&input); err != nil {
		return utils.ResponseError(c, fiber.StatusInternalServerError, "Gagal mendaftarkan user")
	}
	return utils.ResponseCreated(c, nil, "Registrasi berhasil")
}

func Login(c *fiber.Ctx) error {
	var input models.User
	if err := c.BodyParser(&input); err != nil {
		return utils.ResponseError(c, fiber.StatusBadRequest, "Invalid input data")
	}

	user, err := models.GetUserByEmail(input.Email)
	if err != nil {
		return utils.ResponseError(c, fiber.StatusUnauthorized, "Email atau Password salah")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return utils.ResponseError(c, fiber.StatusUnauthorized, "Email atau Password salah")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	t, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	return utils.ResponseSuccess(c, fiber.Map{"token": t}, "Login berhasil")
}
