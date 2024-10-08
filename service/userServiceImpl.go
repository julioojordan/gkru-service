package service

import (
	"database/sql"
	"fmt"
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
	defer helper.CommitOrRollback(tx, logger)

	user, err := service.UserRepository.FindOne(ctx, tx)
	if err != nil {
		fmt.Println("err", err)
		if err.Error() != "user is not found" {
			return nil, fiber.ErrInternalServerError
		}
		return nil, err
	}

	authToken, err := authentication.CreateToken(user.Username)
	helper.PanicIfError(err)

	return helper.ToLoginResponse(authToken), nil
}
