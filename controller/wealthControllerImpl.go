package controller

import (
	"gkru-service/entity"
	"gkru-service/helper"
	"gkru-service/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/sirupsen/logrus"
)

type WealthControllerImpl struct {
	WealthService service.WealthService
}

func NewWealthController(WealthService service.WealthService) WealthController {
	return &WealthControllerImpl{
		WealthService: WealthService,
	}
}

func (controller *WealthControllerImpl) GetTotal(ctx *fiber.Ctx) error {
	logger, _ := ctx.Locals("logger").(*logrus.Logger)
	wealthTotal, err := controller.WealthService.GetTotal(ctx)
	if err != nil {
		return helper.HandleError(ctx, logger, err)
	}
	res := entity.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   wealthTotal,
	}
	logger.WithFields(logrus.Fields{
		"type": "response",
		"code": 200,
		"status": utils.StatusMessage(200),
	}).Info("success")
	return ctx.JSON(res)
}
