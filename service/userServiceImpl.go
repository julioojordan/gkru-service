package service

import (
	"database/sql"
	"gkru-service/authentication"
	"gkru-service/entity"
	"gkru-service/exception"
	"gkru-service/helper"
	"gkru-service/repository"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *UserServiceImpl) FindOne(ctx *fiber.Ctx) entity.LoginResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindOne(ctx, tx)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	// createToken if user found
	authToken, err := authentication.CreateToken(user.Username)
	helper.PanicIfError(err)

	return helper.ToLoginResponse(authToken)
}

