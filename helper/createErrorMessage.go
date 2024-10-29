package helper

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func CreateErrorMessage(baseMessage string, err error) *fiber.Error {
	errorMessage := fmt.Sprintf("%s: %v", baseMessage, err.Error())
	return fiber.NewError(fiber.StatusInternalServerError, errorMessage)
}