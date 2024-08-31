package service

import (
	"database/sql"
	"fmt"
	"gkru-service/helper"
	"gkru-service/repository"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type WealthServiceImpl struct {
	WealthRepository repository.WealthRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewWealthService(wealthRepository repository.WealthRepository, DB *sql.DB, validate *validator.Validate) WealthService {
	return &WealthServiceImpl{
		WealthRepository: wealthRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *WealthServiceImpl) GetTotal(ctx *fiber.Ctx) (interface{}, error) {
	logger, _ := ctx.Locals("logger").(*logrus.Logger)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx, logger)

	totalWealth, err := service.WealthRepository.GetTotal(ctx, tx)
	fmt.Println("total wealth", totalWealth)
	if err != nil {
		return nil, err
	}

	return totalWealth, nil
}
