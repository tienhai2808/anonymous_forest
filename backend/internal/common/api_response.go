package common

import "github.com/gofiber/fiber/v2"

func JSON(c *fiber.Ctx, statusCode int, message string, data any) error {
	return c.Status(statusCode).JSON(ApiResponse{
		Message: message,
		Data:    data,
	})
}
