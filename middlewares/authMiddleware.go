package middlewares

import (
	"fmt"
	"gkru-service/authentication"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func AuthMiddleware(c *fiber.Ctx) error {
	auth := c.Get("Authorization")
	if auth == "" {
		return c.Status(fiber.StatusUnauthorized).SendString("unauthorized")
	}

	tokenString := auth[len("Bearer "):]
	err := authentication.VerifyToken(tokenString)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("unauthorized 2")
	}

	logger, _ := c.Locals("logger").(*logrus.Logger)
	logger.Info("Token valid, request authorized.")

	return c.Next()
}
