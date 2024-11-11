package controller

import (
	"gkru-service/entity"
	"gkru-service/helper"
	"gkru-service/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/sirupsen/logrus"
)

type DataLingkunganControllerImpl struct {
	DataLingkunganService service.DataLingkunganService
}

func NewDataLingkunganController(DataLingkunganService service.DataLingkunganService) DataLingkunganController {
	return &DataLingkunganControllerImpl{
		DataLingkunganService: DataLingkunganService,
	}
}

func (controller *DataLingkunganControllerImpl) FindOneWithParam(ctx *fiber.Ctx) error {
	logger, _ := ctx.Locals("logger").(*logrus.Logger)
	dataLingkungan, err := controller.DataLingkunganService.FindOneWithParam(ctx)
	if err != nil {
		return helper.HandleError(ctx, logger, err)
	}
	res := entity.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   dataLingkungan,
	}
	logger.WithFields(logrus.Fields{
		"type": "response",
		"code": 200,
		"status": utils.StatusMessage(200),
		"data": res.Data,
	}).Info("success")
	return ctx.JSON(res)
}

func (controller *DataLingkunganControllerImpl) FindAll(ctx *fiber.Ctx) error {
	logger, _ := ctx.Locals("logger").(*logrus.Logger)
	dataLingkungan, err := controller.DataLingkunganService.FindAll(ctx)
	if err != nil {
		return helper.HandleError(ctx, logger, err)
	}
	res := entity.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   dataLingkungan,
	}
	logger.WithFields(logrus.Fields{
		"type": "response",
		"code": 200,
		"status": utils.StatusMessage(200),
		"data": res.Data,
	}).Info("success")
	return ctx.JSON(res)
}

func (controller *DataLingkunganControllerImpl) Add(ctx *fiber.Ctx) error {
	logger, _ := ctx.Locals("logger").(*logrus.Logger)
	dataLingkungan, err := controller.DataLingkunganService.Add(ctx)
	if err != nil {
		return helper.HandleError(ctx, logger, err)
	}
	res := entity.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   dataLingkungan,
	}
	logger.WithFields(logrus.Fields{
		"type": "response",
		"code": 200,
		"status": utils.StatusMessage(200),
		"data": res.Data,
	}).Info("success")
	return ctx.JSON(res)
}

func (controller *DataLingkunganControllerImpl) Update(ctx *fiber.Ctx) error {
	logger, _ := ctx.Locals("logger").(*logrus.Logger)
	dataLingkungan, err := controller.DataLingkunganService.Update(ctx)
	if err != nil {
		return helper.HandleError(ctx, logger, err)
	}
	res := entity.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   dataLingkungan,
	}
	logger.WithFields(logrus.Fields{
		"type": "response",
		"code": 200,
		"status": utils.StatusMessage(200),
		"data": res.Data,
	}).Info("success")
	return ctx.JSON(res)
}

func (controller *DataLingkunganControllerImpl) DeleteOne(ctx *fiber.Ctx) error {
	logger, _ := ctx.Locals("logger").(*logrus.Logger)
	dataLingkungan, err := controller.DataLingkunganService.DeleteOne(ctx)
	if err != nil {
		return helper.HandleError(ctx, logger, err)
	}
	res := entity.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   dataLingkungan,
	}
	logger.WithFields(logrus.Fields{
		"type": "response",
		"code": 200,
		"status": utils.StatusMessage(200),
		"data": res.Data,
	}).Info("success")
	return ctx.JSON(res)
}

func (controller *DataLingkunganControllerImpl) GetTotalLingkungan(ctx *fiber.Ctx) error {
	logger, _ := ctx.Locals("logger").(*logrus.Logger)
	dataLingkungan, err := controller.DataLingkunganService.GetTotalLingkungan(ctx)
	if err != nil {
		return helper.HandleError(ctx, logger, err)
	}
	res := entity.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   dataLingkungan,
	}
	logger.WithFields(logrus.Fields{
		"type":   "response",
		"code":   200,
		"status": utils.StatusMessage(200),
		"data":   res.Data,
	}).Info("success")
	return ctx.JSON(res)
}