package middleware

import (
	"fmt"
	"laporan-keuangan/utils"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Protected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return utils.ResponseError(c, fiber.StatusUnauthorized, "Akses ditolak, Token tidak ada")
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			return utils.ResponseError(c, fiber.StatusUnauthorized, "Token tidak valid atau kadaluwarsa")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return utils.ResponseError(c, fiber.StatusUnauthorized, "Token invalid claims")
		}

		c.Locals("user_id", claims["id"])
		return c.Next()
	}
}
