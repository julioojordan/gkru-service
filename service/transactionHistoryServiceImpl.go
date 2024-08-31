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

type TransactionHistoryServiceImpl struct {
	TransactionHistoryRepository repository.TransactionHistoryRepository
	DB                           *sql.DB
	Validate                     *validator.Validate
}

func NewTransactionHistoryService(transactionHistoryRepository repository.TransactionHistoryRepository, DB *sql.DB, validate *validator.Validate) TransactionHistoryService {
	return &TransactionHistoryServiceImpl{
		TransactionHistoryRepository: transactionHistoryRepository,
		DB:                           DB,
		Validate:                     validate,
	}
}

func (service *TransactionHistoryServiceImpl) GetTotalIncome(ctx *fiber.Ctx) (interface{}, error) {
	logger, _ := ctx.Locals("logger").(*logrus.Logger)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx, logger)

	totalWealth, err := service.TransactionHistoryRepository.GetTotalIncome(ctx, tx)
	fmt.Println("total Income", totalWealth)
	if err != nil {
		return nil, err
	}

	return totalWealth, nil
}

func (service *TransactionHistoryServiceImpl) GetTotalOutcome(ctx *fiber.Ctx) (interface{}, error) {
	logger, _ := ctx.Locals("logger").(*logrus.Logger)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx, logger)

	totalWealth, err := service.TransactionHistoryRepository.GetTotalOutcome(ctx, tx)
	fmt.Println("total Outcome", totalWealth)
	if err != nil {
		return nil, err
	}

	return totalWealth, nil
}
