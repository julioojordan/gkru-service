package controller

import (
	"gkru-service/entity"
	"gkru-service/helper"
	"gkru-service/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/sirupsen/logrus"
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
	logger, _ := ctx.Locals("logger").(*logrus.Logger)
	user, err := controller.UserService.FindOne(ctx)
	if err != nil {
		return helper.HandleError(ctx, logger, err)
	}
	res := entity.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   user,
	}
	logger.WithFields(logrus.Fields{
		"type": "response",
		"code": 200,
		"status": utils.StatusMessage(200),
	}).Info("success")
	return ctx.JSON(res)
}

func (controller *UserControllerImpl) FindAll(ctx *fiber.Ctx) error {
	logger, _ := ctx.Locals("logger").(*logrus.Logger)
	dataWilayah, err := controller.UserService.FindAll(ctx)
	if err != nil {
		return helper.HandleError(ctx, logger, err)
	}
	res := entity.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   dataWilayah,
	}
	logger.WithFields(logrus.Fields{
		"type": "response",
		"code": 200,
		"status": utils.StatusMessage(200),
		"data": res.Data,
	}).Info("success")
	return ctx.JSON(res)
}
