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

func (service *DataLingkunganServiceImpl) FindOneById(ctx *fiber.Ctx, id int32) (interface{}, error) {
	logger, _ := ctx.Locals("logger").(*logrus.Logger)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx, logger)

	lingkungan, err := service.DataLingkunganRepository.FindOneById(id, tx)
	if err != nil {
		fmt.Println("err", err)
		if err.Error() != "lingkungan is not found" {
			return nil, fiber.ErrInternalServerError
		}
		return nil, err
	}

	return lingkungan, nil
}
