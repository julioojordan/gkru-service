package middlewares

import (
	"gkru-service/authentication"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func AuthMiddleware(c *fiber.Ctx) error {
	logger, _ := c.Locals("logger").(*logrus.Logger)
	auth := c.Get("Authorization")
	if auth == "" {
		return c.Status(fiber.StatusUnauthorized).SendString("unauthorized")
	}

	tokenString := auth[len("Bearer "):]
	err := authentication.VerifyToken(tokenString, logger)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("unauthorized 2")
	}

	logger.Info("Token valid, request authorized.")

	return c.Next()
}
