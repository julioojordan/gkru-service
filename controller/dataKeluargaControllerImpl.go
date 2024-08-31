package controller

import (
	"gkru-service/entity"
	"gkru-service/helper"
	"gkru-service/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/sirupsen/logrus"
)

type DataKeluargaControllerImpl struct {
	DataKeluargaService service.DataKeluargaService
}

func NewDataKeluargaController(DataKeluargaService service.DataKeluargaService) DataKeluargaController {
	return &DataKeluargaControllerImpl{
		DataKeluargaService: DataKeluargaService,
	}
}

func (controller *DataKeluargaControllerImpl) FindOne(ctx *fiber.Ctx) error {
	logger, _ := ctx.Locals("logger").(*logrus.Logger)
	dataKeluarga, err := controller.DataKeluargaService.FindOne(ctx)
	if err != nil {
		return helper.HandleError(ctx, logger, err)
	}
	res := entity.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   dataKeluarga,
	}
	logger.WithFields(logrus.Fields{
		"type": "response",
		"code": 200,
		"status": utils.StatusMessage(200),
		"data": res.Data,
	}).Info("success")
	return ctx.JSON(res)
}

func (controller *DataKeluargaControllerImpl) AddKeluarga(ctx *fiber.Ctx) error {
	logger, _ := ctx.Locals("logger").(*logrus.Logger)
	dataKeluarga, err := controller.DataKeluargaService.AddKeluarga(ctx)
	if err != nil {
		return helper.HandleError(ctx, logger, err)
	}
	res := entity.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   dataKeluarga,
	}
	logger.WithFields(logrus.Fields{
		"type": "response",
		"code": 200,
		"status": utils.StatusMessage(200),
		"data": res.Data,
	}).Info("success")
	return ctx.JSON(res)
}

func (controller *DataKeluargaControllerImpl) GetTotalKeluarga(ctx *fiber.Ctx) error {
	logger, _ := ctx.Locals("logger").(*logrus.Logger)
	dataKeluarga, err := controller.DataKeluargaService.GetTotalKeluarga(ctx)
	if err != nil {
		return helper.HandleError(ctx, logger, err)
	}
	res := entity.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   dataKeluarga,
	}
	logger.WithFields(logrus.Fields{
		"type": "response",
		"code": 200,
		"status": utils.StatusMessage(200),
		"data": res.Data,
	}).Info("success")
	return ctx.JSON(res)
}

func (controller *DataKeluargaControllerImpl) UpdateDataKeluarga(ctx *fiber.Ctx) error {
	logger, _ := ctx.Locals("logger").(*logrus.Logger)
	dataKeluarga, err := controller.DataKeluargaService.UpdateDataKeluarga(ctx)
	if err != nil {
		return helper.HandleError(ctx, logger, err)
	}
	res := entity.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   dataKeluarga,
	}
	logger.WithFields(logrus.Fields{
		"type": "response",
		"code": 200,
		"status": utils.StatusMessage(200),
		"data": res.Data,
	}).Info("success")
	return ctx.JSON(res)
}
