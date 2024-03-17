package controller

import (
	"gkru-service/entity"
	"gkru-service/service"

	"github.com/gofiber/fiber/v2"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) FindOne(ctx *fiber.Ctx) error {
	user := controller.UserService.FindOne(ctx)
	res := entity.WebResponse{
		Code: 200,
		Status: "Ok",
		Data: user,
	}
	return ctx.JSON(res)
}