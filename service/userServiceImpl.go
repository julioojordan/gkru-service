package service

import (
	"database/sql"
	"gkru-service/authentication"
	"gkru-service/helper"
	"gkru-service/repository"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *UserServiceImpl) FindOne(ctx *fiber.Ctx) (interface{}, error) {
	logger, _ := ctx.Locals("logger").(*logrus.Logger)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer func() {
		helper.CommitOrRollback2(tx, logger, err) // Selalu panggil CommitOrRollback2
	}()

	user, err := service.UserRepository.FindOne(ctx, tx)
	if err != nil {
		return nil, err
	}

	authToken, err := authentication.CreateToken(user.Username)
	helper.PanicIfError(err)

	return helper.ToLoginResponse(authToken, user), nil
}

func (service *UserServiceImpl) FindAll(ctx *fiber.Ctx) (interface{}, error) {
	logger, _ := ctx.Locals("logger").(*logrus.Logger)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer func() {
		helper.CommitOrRollback2(tx, logger, err) // Selalu panggil CommitOrRollback2
	}()

	result, err := service.UserRepository.FindAll(ctx, tx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (service *UserServiceImpl) Update(ctx *fiber.Ctx) (interface{}, error) {
	logger, _ := ctx.Locals("logger").(*logrus.Logger)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer func() {
		helper.CommitOrRollback2(tx, logger, err) // Selalu panggil CommitOrRollback2
	}()

	result, err := service.UserRepository.Update(ctx, tx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (service *UserServiceImpl) Add(ctx *fiber.Ctx) (interface{}, error) {
	logger, _ := ctx.Locals("logger").(*logrus.Logger)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer func() {
		helper.CommitOrRollback2(tx, logger, err) // Selalu panggil CommitOrRollback2
	}()

	result, err := service.UserRepository.Add(ctx, tx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (service *UserServiceImpl) DeleteOne(ctx *fiber.Ctx) (interface{}, error) {
	logger, _ := ctx.Locals("logger").(*logrus.Logger)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer func() {
		helper.CommitOrRollback2(tx, logger, err) // Selalu panggil CommitOrRollback2
	}()

	result, err := service.UserRepository.DeleteOne(ctx, tx)
	if err != nil {
		return nil, err
	}

	return result, nil
}
