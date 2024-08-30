package service

import (
	"database/sql"
	"gkru-service/helper"
	"gkru-service/repository"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type DataAnggotaServiceImpl struct {
	DataAnggotaRepository repository.DataAnggotaRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewDataAnggotaService(dataAnggotaRepository repository.DataAnggotaRepository, DB *sql.DB, validate *validator.Validate) DataAnggotaService {
	return &DataAnggotaServiceImpl{
		DataAnggotaRepository: dataAnggotaRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *DataAnggotaServiceImpl) AddAnggota(ctx *fiber.Ctx) (interface{}, error) {
	logger, _ := ctx.Locals("logger").(*logrus.Logger)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx, logger)

	result, err := service.DataAnggotaRepository.AddAnggota(ctx, tx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (service *DataAnggotaServiceImpl) GetTotalAnggota(ctx *fiber.Ctx) (interface{}, error) {
	logger, _ := ctx.Locals("logger").(*logrus.Logger)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx, logger)

	result, err := service.DataAnggotaRepository.GetTotalAnggota(ctx, tx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (service *DataAnggotaServiceImpl) UpdateAnggota(ctx *fiber.Ctx) (interface{}, error) {
	logger, _ := ctx.Locals("logger").(*logrus.Logger)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx, logger)

	result, err := service.DataAnggotaRepository.UpdateAnggota(ctx, tx)
	if err != nil {
		return nil, err
	}

	return result, nil
}