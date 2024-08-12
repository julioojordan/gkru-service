package controller

import (
	"gkru-service/entity"
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
		Data:   dataKeluarga,
	}
	logger.WithFields(logrus.Fields{
		"type": "response",
		"code": 200,
		"status": utils.StatusMessage(200),
	}).Info("success")
	return ctx.JSON(res)
}
