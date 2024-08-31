package helper

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/sirupsen/logrus"
)

func HandleError(ctx *fiber.Ctx, logger *logrus.Logger, err error) error {
	if res, ok := err.(*fiber.Error); ok {
		logger.WithFields(logrus.Fields{
			"type":   "response",
			"code":   res.Code,
			"status": utils.StatusMessage(res.Code),
		}).Warn(res.Error())

		return ctx.Status(res.Code).JSON(fiber.Map{
			"code":    res.Code,
			"status":  utils.StatusMessage(res.Code),
			"message": res.Message,
		})
	}else {
		logger.WithFields(logrus.Fields{
			"type":   "response",
			"code":   fiber.StatusInternalServerError,
			"status": "Internal Server Error",
		}).Error(err.Error())
	
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"status":  "Internal Server Error",
			"message": "Something went wrong",
		})
	}
}