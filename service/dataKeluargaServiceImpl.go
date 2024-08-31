package service

import (
	"database/sql"
	"gkru-service/helper"
	"gkru-service/repository"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type DataKeluargaServiceImpl struct {
	DataKeluargaRepository repository.DataKeluargaRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewDataKeluargaService(DataKeluargaRepository repository.DataKeluargaRepository, DB *sql.DB, validate *validator.Validate) DataKeluargaService {
	return &DataKeluargaServiceImpl{
		DataKeluargaRepository: DataKeluargaRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *DataKeluargaServiceImpl) FindOne(ctx *fiber.Ctx) (interface{}, error) {
	logger, _ := ctx.Locals("logger").(*logrus.Logger)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx, logger)

	dataKeluarga, err := service.DataKeluargaRepository.FindOne(ctx, tx)
	if err != nil {
		return nil, err
	}

	return dataKeluarga, nil
}

func (service *DataKeluargaServiceImpl) AddKeluarga(ctx *fiber.Ctx) (interface{}, error) {
	logger, _ := ctx.Locals("logger").(*logrus.Logger)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx, logger)

	result, err := service.DataKeluargaRepository.AddKeluarga(ctx, tx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (service *DataKeluargaServiceImpl) GetTotalKeluarga(ctx *fiber.Ctx) (interface{}, error) {
	logger, _ := ctx.Locals("logger").(*logrus.Logger)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx, logger)

	result, err := service.DataKeluargaRepository.GetTotalKeluarga(ctx, tx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (service *DataKeluargaServiceImpl) UpdateDataKeluarga(ctx *fiber.Ctx) (interface{}, error) {
	logger, _ := ctx.Locals("logger").(*logrus.Logger)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx, logger)

	result, err := service.DataKeluargaRepository.UpdateDataKeluarga(ctx, tx)
	if err != nil {
		return nil, err
	}

	return result, nil
}
