package controller

import "github.com/gofiber/fiber/v2"

type UserController interface {
	FindOne(ctx *fiber.Ctx) error
}