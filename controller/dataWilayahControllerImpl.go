package controller

import (
	"gkru-service/entity"
	"gkru-service/helper"
	"gkru-service/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/sirupsen/logrus"
)

type DataWilayahControllerImpl struct {
	DataWilayahService service.DataWilayahService
}

func NewDataWilayahController(DataWilayahService service.DataWilayahService) DataWilayahController {
	return &DataWilayahControllerImpl{
		DataWilayahService: DataWilayahService,
	}
}

func (controller *DataWilayahControllerImpl) FindOne(ctx *fiber.Ctx) error {
	logger, _ := ctx.Locals("logger").(*logrus.Logger)
	dataWilayah, err := controller.DataWilayahService.FindOne(ctx)
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

func (controller *DataWilayahControllerImpl) FindAll(ctx *fiber.Ctx) error {
	logger, _ := ctx.Locals("logger").(*logrus.Logger)
	dataWilayah, err := controller.DataWilayahService.FindAll(ctx)
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

func (controller *DataWilayahControllerImpl) Add(ctx *fiber.Ctx) error {
	logger, _ := ctx.Locals("logger").(*logrus.Logger)
	dataWilayah, err := controller.DataWilayahService.Add(ctx)
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

func (controller *DataWilayahControllerImpl) Update(ctx *fiber.Ctx) error {
	logger, _ := ctx.Locals("logger").(*logrus.Logger)
	dataWilayah, err := controller.DataWilayahService.Update(ctx)
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

func (controller *DataWilayahControllerImpl) DeleteOne(ctx *fiber.Ctx) error {
	logger, _ := ctx.Locals("logger").(*logrus.Logger)
	dataWilayah, err := controller.DataWilayahService.DeleteOne(ctx)
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

func (controller *DataWilayahControllerImpl) GetTotalWilayah(ctx *fiber.Ctx) error {
	logger, _ := ctx.Locals("logger").(*logrus.Logger)
	dataWilayah, err := controller.DataWilayahService.GetTotalWilayah(ctx)
	if err != nil {
		return helper.HandleError(ctx, logger, err)
	}
	res := entity.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   dataWilayah,
	}
	logger.WithFields(logrus.Fields{
		"type":   "response",
		"code":   200,
		"status": utils.StatusMessage(200),
		"data":   res.Data,
	}).Info("success")
	return ctx.JSON(res)
}