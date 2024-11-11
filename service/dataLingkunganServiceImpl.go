package service

import (
	"database/sql"
	"gkru-service/helper"
	"gkru-service/repository"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type DataLingkunganServiceImpl struct {
	DataLingkunganRepository repository.DataLingkunganRepository
	DB                       *sql.DB
	Validate                 *validator.Validate
}

func NewDataLingkunganService(DataLingkunganRepository repository.DataLingkunganRepository, DB *sql.DB, validate *validator.Validate) DataLingkunganService {
	return &DataLingkunganServiceImpl{
		DataLingkunganRepository: DataLingkunganRepository,
		DB:                       DB,
		Validate:                 validate,
	}
}

func (service *DataLingkunganServiceImpl) FindOneWithParam(ctx *fiber.Ctx) (interface{}, error) {
	logger, _ := ctx.Locals("logger").(*logrus.Logger)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx, logger)

	result, err := service.DataLingkunganRepository.FindOneWithParam(ctx, tx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (service *DataLingkunganServiceImpl) FindAll(ctx *fiber.Ctx) (interface{}, error) {
	logger, _ := ctx.Locals("logger").(*logrus.Logger)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx, logger)

	result, err := service.DataLingkunganRepository.FindAll(ctx, tx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (service *DataLingkunganServiceImpl) Add(ctx *fiber.Ctx) (interface{}, error) {
	logger, _ := ctx.Locals("logger").(*logrus.Logger)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx, logger)

	result, err := service.DataLingkunganRepository.Add(ctx, tx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (service *DataLingkunganServiceImpl) Update(ctx *fiber.Ctx) (interface{}, error) {
	logger, _ := ctx.Locals("logger").(*logrus.Logger)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx, logger)

	result, err := service.DataLingkunganRepository.Update(ctx, tx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (service *DataLingkunganServiceImpl) DeleteOne(ctx *fiber.Ctx) (interface{}, error) {
	logger, _ := ctx.Locals("logger").(*logrus.Logger)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx, logger)

	result, err := service.DataLingkunganRepository.DeleteOne(ctx, tx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (service *DataLingkunganServiceImpl) GetTotalLingkungan(ctx *fiber.Ctx) (interface{}, error) {
	logger, _ := ctx.Locals("logger").(*logrus.Logger)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx, logger)

	result, err := service.DataLingkunganRepository.GetTotalLingkungan(ctx, tx)
	if err != nil {
		return nil, err
	}

	return result, nil
}
