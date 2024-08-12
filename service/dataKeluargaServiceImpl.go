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
	fmt.Println("dataKeluarga", dataKeluarga)
	if err != nil {
		fmt.Println("err", err)
		if err.Error() != "Data Keluarga is not found" {
			return nil, fiber.ErrInternalServerError
		}
		return nil, err
	}

	return dataKeluarga, nil
}
