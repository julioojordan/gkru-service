package controller

import (
	"gkru-service/entity"
	"gkru-service/helper"
	"gkru-service/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/sirupsen/logrus"
)

type DataAnggotaControllerImpl struct {
	DataAnggotaService service.DataAnggotaService
}

func NewDataAnggotaController(DataAnggotaService service.DataAnggotaService) DataAnggotaController {
	return &DataAnggotaControllerImpl{
		DataAnggotaService: DataAnggotaService,
	}
}

func (controller *DataAnggotaControllerImpl) AddAnggota(ctx *fiber.Ctx) error {
	logger, _ := ctx.Locals("logger").(*logrus.Logger)
	dataAnggota, err := controller.DataAnggotaService.AddAnggota(ctx)
	if err != nil {
		return helper.HandleError(ctx, logger, err)
	}
	res := entity.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   dataAnggota,
	}
	logger.WithFields(logrus.Fields{
		"type": "response",
		"code": 200,
		"status": utils.StatusMessage(200),
	}).Info("success")
	return ctx.JSON(res)
}

func (controller *DataAnggotaControllerImpl) UpdateAnggota(ctx *fiber.Ctx) error {
	logger, _ := ctx.Locals("logger").(*logrus.Logger)
	dataAnggota, err := controller.DataAnggotaService.UpdateAnggota(ctx)
	if err != nil {
		return helper.HandleError(ctx, logger, err)
	}
	res := entity.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   dataAnggota,
	}
	logger.WithFields(logrus.Fields{
		"type": "response",
		"code": 200,
		"status": utils.StatusMessage(200),
	}).Info("success")
	return ctx.JSON(res)
}

func (controller *DataAnggotaControllerImpl) GetTotalAnggota(ctx *fiber.Ctx) error {
	logger, _ := ctx.Locals("logger").(*logrus.Logger)
	dataAnggota, err := controller.DataAnggotaService.GetTotalAnggota(ctx)
	if err != nil {
		return helper.HandleError(ctx, logger, err)
	}
	res := entity.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   dataAnggota,
	}
	logger.WithFields(logrus.Fields{
		"type": "response",
		"code": 200,
		"status": utils.StatusMessage(200),
	}).Info("success")
	return ctx.JSON(res)
}
