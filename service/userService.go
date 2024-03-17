package service

import (
	"gkru-service/entity"

	"github.com/gofiber/fiber/v2"
)

type UserService interface {
	FindOne(ctx *fiber.Ctx) entity.UserResponse
}