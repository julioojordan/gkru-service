package service

import (
	"database/sql"
	"gkru-service/helper"
	"gkru-service/repository"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type DataWilayahServiceImpl struct {
	DataWilayahRepository repository.DataWilayahRepository
	DB                    *sql.DB
	Validate              *validator.Validate
}

func NewDataWilayahService(DataWilayahRepository repository.DataWilayahRepository, DB *sql.DB, validate *validator.Validate) DataWilayahService {
	return &DataWilayahServiceImpl{
		DataWilayahRepository: DataWilayahRepository,
		DB:                    DB,
		Validate:              validate,
	}
}

func (service *DataWilayahServiceImpl) FindOne(ctx *fiber.Ctx) (interface{}, error) {
	logger, _ := ctx.Locals("logger").(*logrus.Logger)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx, logger)

	result, err := service.DataWilayahRepository.FindOne(ctx, tx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (service *DataWilayahServiceImpl) FindAll(ctx *fiber.Ctx) (interface{}, error) {
	logger, _ := ctx.Locals("logger").(*logrus.Logger)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx, logger)

	result, err := service.DataWilayahRepository.FindAll(ctx, tx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (service *DataWilayahServiceImpl) Add(ctx *fiber.Ctx) (interface{}, error) {
	logger, _ := ctx.Locals("logger").(*logrus.Logger)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx, logger)

	result, err := service.DataWilayahRepository.Add(ctx, tx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (service *DataWilayahServiceImpl) Update(ctx *fiber.Ctx) (interface{}, error) {
	logger, _ := ctx.Locals("logger").(*logrus.Logger)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx, logger)

	result, err := service.DataWilayahRepository.Update(ctx, tx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (service *DataWilayahServiceImpl) DeleteOne(ctx *fiber.Ctx) (interface{}, error) {
	logger, _ := ctx.Locals("logger").(*logrus.Logger)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx, logger)

	result, err := service.DataWilayahRepository.DeleteOne(ctx, tx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (service *DataWilayahServiceImpl) GetTotalWilayah(ctx *fiber.Ctx) (interface{}, error) {
	logger, _ := ctx.Locals("logger").(*logrus.Logger)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx, logger)

	result, err := service.DataWilayahRepository.GetTotalWilayah(ctx, tx)
	if err != nil {
		return nil, err
	}

	return result, nil
}