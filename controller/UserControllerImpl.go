package controller

import (
	"gkru-service/entity"
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
		// manually type checking
		if res, ok := err.(*fiber.Error); ok {
			logger.WithFields(logrus.Fields{
				"type": "response",
				"code": res.Code,
				"status": utils.StatusMessage(res.Code),
			}).Warn(res.Error())

			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"code":    res.Code,
				"status": utils.StatusMessage(res.Code),
				"message": res.Error(),
			})
		}
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
