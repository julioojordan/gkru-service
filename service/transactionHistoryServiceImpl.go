package service

import (
	"database/sql"
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
	if err != nil {
		return nil, err
	}

	return totalWealth, nil
}

func (service *TransactionHistoryServiceImpl) FindAll(ctx *fiber.Ctx) (interface{}, error) {
	logger, _ := ctx.Locals("logger").(*logrus.Logger)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx, logger)

	totalWealth, err := service.TransactionHistoryRepository.FindAll(ctx, tx)
	if err != nil {
		return nil, err
	}

	return totalWealth, nil
}

func (service *TransactionHistoryServiceImpl) FindAllWithKeluargaContext(ctx *fiber.Ctx) (interface{}, error) {
	logger, _ := ctx.Locals("logger").(*logrus.Logger)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx, logger)

	totalWealth, err := service.TransactionHistoryRepository.FindAllWithKeluargaContext(ctx, tx)
	if err != nil {
		return nil, err
	}

	return totalWealth, nil
}

func (service *TransactionHistoryServiceImpl) FindOne(ctx *fiber.Ctx) (interface{}, error) {
	logger, _ := ctx.Locals("logger").(*logrus.Logger)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx, logger)

	totalWealth, err := service.TransactionHistoryRepository.FindOne(ctx, tx)
	if err != nil {
		return nil, err
	}

	return totalWealth, nil
}

func (service *TransactionHistoryServiceImpl) FindAllWithIdKeluarga(ctx *fiber.Ctx) (interface{}, error) {
	logger, _ := ctx.Locals("logger").(*logrus.Logger)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx, logger)

	totalWealth, err := service.TransactionHistoryRepository.FindAllWithIdKeluarga(ctx, tx)
	if err != nil {
		return nil, err
	}

	return totalWealth, nil
}

func (service *TransactionHistoryServiceImpl) Update(ctx *fiber.Ctx) (interface{}, error) {
	logger, _ := ctx.Locals("logger").(*logrus.Logger)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx, logger)

	totalWealth, err := service.TransactionHistoryRepository.Update(ctx, tx)
	if err != nil {
		return nil, err
	}

	return totalWealth, nil
}

func (service *TransactionHistoryServiceImpl) Delete(ctx *fiber.Ctx) (interface{}, error) {
	logger, _ := ctx.Locals("logger").(*logrus.Logger)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx, logger)

	totalWealth, err := service.TransactionHistoryRepository.Delete(ctx, tx)
	if err != nil {
		return nil, err
	}

	return totalWealth, nil
}

func (service *TransactionHistoryServiceImpl) Add(ctx *fiber.Ctx) (interface{}, error) {
	logger, _ := ctx.Locals("logger").(*logrus.Logger)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx, logger)

	totalWealth, err := service.TransactionHistoryRepository.Add(ctx, tx)
	if err != nil {
		return nil, err
	}

	return totalWealth, nil
}