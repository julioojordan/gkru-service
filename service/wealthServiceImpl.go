package service

import (
	"database/sql"
	"gkru-service/helper"
	"gkru-service/repository"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type WealthServiceImpl struct {
	WealthRepository repository.WealthRepository
	DB               *sql.DB
	Validate         *validator.Validate
}

func NewWealthService(wealthRepository repository.WealthRepository, DB *sql.DB, validate *validator.Validate) WealthService {
	return &WealthServiceImpl{
		WealthRepository: wealthRepository,
		DB:               DB,
		Validate:         validate,
	}
}

func (service *WealthServiceImpl) GetTotal(ctx *fiber.Ctx) (interface{}, error) {
	logger, _ := ctx.Locals("logger").(*logrus.Logger)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer func() {
		helper.CommitOrRollback2(tx, logger, err) // Selalu panggil CommitOrRollback2
	}()

	totalWealth, err := service.WealthRepository.GetTotal(ctx, tx)
	if err != nil {
		return nil, err
	}

	return totalWealth, nil
}
