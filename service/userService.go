package service

import (

	"github.com/gofiber/fiber/v2"
)

type UserService interface {
	FindOne(ctx *fiber.Ctx) (interface{}, error)
}