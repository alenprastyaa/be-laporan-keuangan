package utils

import "github.com/gofiber/fiber/v2"

type ApiResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func ResponseSuccess(c *fiber.Ctx, data interface{}, message string) error {
	return c.Status(fiber.StatusOK).JSON(ApiResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func ResponseCreated(c *fiber.Ctx, data interface{}, message string) error {
	return c.Status(fiber.StatusCreated).JSON(ApiResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func ResponseError(c *fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(ApiResponse{
		Success: false,
		Message: message,
		Data:    nil,
	})
}
