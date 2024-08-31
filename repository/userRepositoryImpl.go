package repository

import (
	"database/sql"
	"encoding/json"
	"gkru-service/entity"
	"gkru-service/helper"

	"github.com/gofiber/fiber/v2"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type userRepositoryImpl struct {
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepositoryImpl{}
}

func (repository *userRepositoryImpl) FindOne(ctx *fiber.Ctx, tx *sql.Tx) (entity.User, error) {
	sqlScript := "SELECT id, username FROM users WHERE username = ? AND password = ?"
	body := ctx.Body()
	request := new(LoginRequest)
	err := json.Unmarshal(body, request)
	helper.PanicIfError(err)

	result, err :=tx.Query(sqlScript, request.Username, request.Password)
	helper.PanicIfError(err);
	defer result.Close()
	
	user := entity.User{}
	if result.Next(){
		err := result.Scan(&user.Id, &user.Username)
		helper.PanicIfError(err)
		return user, nil
	} else{
		return user, fiber.NewError(fiber.StatusNotFound, "user is not found")
	}
}
