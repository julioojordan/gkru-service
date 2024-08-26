package controller

import (
	"gkru-service/entity"
	"gkru-service/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/sirupsen/logrus"
)

type TransactionHistoryControllerImpl struct {
	TransactionHistoryService service.TransactionHistoryService
}

func NewTransactionHistoryController(transactionHistoryService service.TransactionHistoryService) TransactionHistoryController {
	return &TransactionHistoryControllerImpl{
		TransactionHistoryService: transactionHistoryService,
	}
}

func (controller *TransactionHistoryControllerImpl) GetTotalIncome(ctx *fiber.Ctx) error {
	logger, _ := ctx.Locals("logger").(*logrus.Logger)
	totalIncome, err := controller.TransactionHistoryService.GetTotalIncome(ctx)
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
		Data:   totalIncome,
	}
	logger.WithFields(logrus.Fields{
		"type": "response",
		"code": 200,
		"status": utils.StatusMessage(200),
	}).Info("success")
	return ctx.JSON(res)
}

func (controller *TransactionHistoryControllerImpl) GetTotalOutcome(ctx *fiber.Ctx) error {
	logger, _ := ctx.Locals("logger").(*logrus.Logger)
	totalOutcome, err := controller.TransactionHistoryService.GetTotalOutcome(ctx)
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
		Data:   totalOutcome,
	}
	logger.WithFields(logrus.Fields{
		"type": "response",
		"code": 200,
		"status": utils.StatusMessage(200),
	}).Info("success")
	return ctx.JSON(res)
}
